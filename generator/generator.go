package generator

import (
	"github.com/elahe-dastan/urlShortener_KGS/model"
	"github.com/jinzhu/gorm"
)

const LengthOfShortURL = 8

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

func Generate(db *gorm.DB) {

	generateURLsRec("", LengthOfShortURL, db)
}
