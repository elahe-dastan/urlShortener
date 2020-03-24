package generator

import (
	"database/sql"
	"log"
)

const source = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateURLsRec(prefix string, k int, db *sql.DB) {
	if k == 0 {
		_, err := db.Exec("INSERT INTO short_url VALUES ($1, $2)", prefix, false)
		if err != nil {
			log.Printf("Cannot insert row of KGS due to the following error\n %s", err)
		}

		return
	}

	k--

	for i := range source {
		newPrefix := prefix + string(source[i])
		generateURLsRec(newPrefix, k, db)
	}
}

func Generate(db *sql.DB, l int) {
	generateURLsRec("", l, db)
}
