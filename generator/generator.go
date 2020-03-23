package generator

import (
	"database/sql"
	"fmt"
	"log"
)

const source = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateURLsRec(prefix string, k int, db *sql.DB) {
	if k == 0 {
		_, err := db.Exec("INSERT INTO short_url VALUES ($1, $2)", prefix, false)
		if err != nil {
			log.Fatal(err)
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
	fmt.Println("Length of short URL")
	fmt.Println(l)
	generateURLsRec("", l, db)
}
