package request

import (
	"time"

	"github.com/elahe-dastan/urlShortener/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Map struct {
	LongURL        string
	ShortURL       string
	ExpirationTime time.Time
}

func (r Map) Validate() bool {
	err := validation.ValidateStruct(&r,
		validation.Field(&r.LongURL, validation.Required, is.URL))

	return err == nil
}

func (r Map) Model() model.Map {
	var m model.Map

	m.ShortURL = r.ShortURL
	m.LongURL = r.LongURL
	m.ExpirationTime = r.ExpirationTime

	return m
}
