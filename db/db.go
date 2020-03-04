package db

import (
	"log"

	"github.com/elahe-dastan/urlShortener_KGS/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //adding dialect for postgres
)

func New(config config.Database) *gorm.DB {
	db, err := gorm.Open(config.DBName, config.ConnectionString)
	if err != nil {
		log.Fatalf("can not open connection to database due to the following err\n: %s", err)
	}

	return db
}
