package model

import "time"

type Map struct {
	ID             int    `gorm:"primary_key"`
	LongURL        string `gorm:"not null"`
	ShortURL       string `gorm:"unique;not null"`
	ExpirationTime time.Time
}
