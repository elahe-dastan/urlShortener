package models

type RandomShortURL struct {
	ShortURL  string   `gorm:"primary_key"`
	IsUsed bool
}
