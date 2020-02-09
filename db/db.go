package db

import (
	"bufio"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"urlShortener/KGS"
	"urlShortener/models"
)


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

	//Begin a transaction so two threads cannot change a row at the same time
	tx := db.Begin()

	if tx.Error != nil {
		fmt.Print(tx.Error)
	}

	var selectedShortURL models.RandomShortURL
	tx.Where("is_used = ?", false).First(&selectedShortURL)

	selectedShortURL.IsUsed = true
	tx.Save(&selectedShortURL)

	if err := tx.Commit().Error; err != nil {
		fmt.Printf("There is an error: %s\n", err.Error())
	}
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

	db.Raw("SELECT * from short_url_maps WHERE short_url = ?;", url).Scan(&mapping)
	return mapping
}

func connectToDatabase() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=koochooloo password=postgres")
	if err != nil {
		fmt.Errorf("db new client error: %s", err)
	}
	return db
}
