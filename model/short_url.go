package model

type ShortURL struct {
	URL    string `gorm:"primary_key"`
	IsUsed bool
}
