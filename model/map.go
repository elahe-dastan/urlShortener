package model

import "time"

type Map struct {
	Id   int  `gorm:"primary_key"`
	LongURL  string `gorm:"not null"`
	ShortURL string  `gorm:"unique;not null"`
	ExpirationTime time.Time
}