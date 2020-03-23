package db

import (
	"database/sql"
	"log"

	"github.com/elahe-dastan/urlShortener/config"
	_ "github.com/lib/pq" //adding dialect for postgres
)

const DB = "postgres"

func New(config config.Database) *sql.DB {
	db, err := sql.Open(DB, config.Cstring())
	if err != nil {
		log.Fatalf("can not open connection to database due to the following err\n: %s", err)
	}

	return db
}
