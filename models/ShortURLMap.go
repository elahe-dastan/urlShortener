package models

type ShortURLMap struct {
	Id   int  `gorm:"primary_key"`
	LongURL  string `gorm:"unique;not null"`
	ShortURL string  `gorm:"not null"`
}