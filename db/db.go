package db

import (
	"bufio"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"urlShortener/KGS"
	"urlShortener/configuration"
	"urlShortener/models"
)


var config configuration.Constants

func SetConfiguration(constants configuration.Constants)  {
	config = constants
}

func SaveRandomShortURLs() {
	var db = connectToDatabase()

	if db.HasTable(&models.RandomShortURL{}) {
		return
	}

	randomShortURLs := KGS.GenerateAllRandomShortURLs()
	db.Debug().AutoMigrate(&models.RandomShortURL{})

	for _, shortUrl := range randomShortURLs {
		db.Create(&shortUrl)
	}

	defer db.Close()
}

func CreateMapTable() {
	var db = connectToDatabase()

	if db.HasTable(&models.RandomShortURL{}) {
		return
	}

	db.Debug().AutoMigrate(&models.ShortURLMap{})

	defer db.Close()
}

func ChooseShortURLTransaction()  string {
	var db = connectToDatabase()

	var selectedShortURL models.RandomShortURL
	db.Raw("UPDATE random_short_urls SET is_used = ? WHERE short_url = " +
		"(SELECT short_url FROM random_short_urls WHERE is_used = ? LIMIT 1) " +
		"RETURNING *;", true, false).Scan(&selectedShortURL)
	return selectedShortURL.ShortURL
}

func ChooseShortURLRowLock()  string {
	var db = connectToDatabase()

	var selectedShortURL models.RandomShortURL
	db.Raw("SELECT * FROM random_short_urls WHERE is_used = ? LIMIT = 1 FOR UPDATE", false).Scan(&selectedShortURL)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	selectedShortURL.IsUsed = true
	db.Save(&selectedShortURL)

	db.Exec("COMMIT")
	return selectedShortURL.ShortURL
}

func InsertToMapping(urlMap models.ShortURLMap)  {
	var db = connectToDatabase()
	db.Create(&urlMap)
}

func RetriveLongURL(url string) models.ShortURLMap {
	var db = connectToDatabase()
	var mapping models.ShortURLMap

	db.Raw("SELECT * from short_url_maps WHERE short_url = ?;", url).Scan(&mapping) //O(lgn)
	return mapping
}

func connectToDatabase() *gorm.DB {
	db, err := gorm.Open(config.DatabaseConfig.DBName, config.DatabaseConfig.ConnectionString)
	if err != nil {
		fmt.Errorf("db new client error: %s", err)
	}
	return db
}


