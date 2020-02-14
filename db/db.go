package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
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

	if db.HasTable(&models.ShortToLongURLMap{}) {
		return
	}

	db.Debug().AutoMigrate(&models.ShortToLongURLMap{})

	db.Exec("create or replace function delete_expired_row() " +
					"returns trigger as " +
					"$BODY$ " +
					"begin " +
					"delete from short_to_long_url_maps where expiration_time < NOW(); " +
					"return null; " +
					"end; " +
					"$BODY$ " +
					"LANGUAGE plpgsql;" +
				"create trigger delete_expired_rows " +
					"after insert " +
					"on short_to_long_url_maps " +
					"for each row " +
					"execute procedure delete_expired_row();")

	defer db.Close()
}

func ChooseShortURLTransaction() string {
	var db = connectToDatabase()
	defer db.Close()

	var selectedShortURL models.RandomShortURL
	db.Raw("UPDATE random_short_urls SET is_used = ? WHERE short_url = " +
		"(SELECT short_url FROM random_short_urls WHERE is_used = ? LIMIT 1) " +
		"RETURNING *;", true, false).Scan(&selectedShortURL) //O(lgn)
	return selectedShortURL.ShortURL
}

func InsertToMapping(urlMap models.ShortToLongURLMap) error {
	var db = connectToDatabase()
	err := db.Create(&urlMap).Error
	return err
}

func RetrieveLongURL(url string) (models.ShortToLongURLMap, error) {
	var db = connectToDatabase()
	var mapping models.ShortToLongURLMap

	db.Raw("SELECT * from short_url_maps WHERE short_url = ?;", url).Scan(&mapping) //O(lgn)

	var err error
	if mapping.LongURL == "" {
		err = errors.New("This short URL doesn't exist in the database")
	}

	return mapping, err
}

func connectToDatabase() *gorm.DB {
	db, err := gorm.Open(config.DatabaseConfig.DBName, config.DatabaseConfig.ConnectionString)
	if err != nil {
		log.Fatalf("can not open connection to datbase due to the following err\n: %s", err)
	}
	return db
}