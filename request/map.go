package request

import (
	"net/url"
	"time"

	"github.com/elahe-dastan/urlShortener_KGS/model"
)

type Map struct {
	LongURL        string
	ShortURL       string
	ExpirationTime time.Time
}

func (r Map) CheckLongURL() bool {
	_, err := url.ParseRequestURI(r.LongURL)

	return err == nil
}

func (r Map) CreateModel() model.Map {
	var m model.Map

	m.ShortURL = r.ShortURL
	m.LongURL = r.LongURL
	m.ExpirationTime = r.ExpirationTime

	return m
}
