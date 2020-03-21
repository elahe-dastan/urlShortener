package generator

import (
	"fmt"

	"github.com/elahe-dastan/urlShortener/config"
	"github.com/elahe-dastan/urlShortener/model"
	"github.com/jinzhu/gorm"
)

const source = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateURLsRec(prefix string, k int, db *gorm.DB) {
	if k == 0 {
		db.Create(&model.ShortURL{URL: prefix, IsUsed: false})
		return
	}

	k--

	for i := range source {
		newPrefix := prefix + string(source[i])
		generateURLsRec(newPrefix, k, db)
	}
}

func Generate(db *gorm.DB, u config.ShortURL) {
	fmt.Println("Length of short URL")
	fmt.Println(u.Length)
	generateURLsRec("", u.Length, db)
}
