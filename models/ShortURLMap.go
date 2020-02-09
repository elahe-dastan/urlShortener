package models

type ShortURLMap struct {
	Id   int  `gorm:"primary_key"`
	LongURL  string `gorm:"not null"`
	ShortURL string  `gorm:"unique;not null"`
}