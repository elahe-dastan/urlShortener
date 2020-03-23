package model

import "time"

type Map struct {
	ID             int
	LongURL        string
	ShortURL       string
	ExpirationTime time.Time
}
