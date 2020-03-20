package main

import (
	"testing"
	"time"

	"github.com/elahe-dastan/urlShortener/config"
	"github.com/elahe-dastan/urlShortener/db"
	"github.com/elahe-dastan/urlShortener/model"
	"github.com/elahe-dastan/urlShortener/store"
)

//Choose a random row and check if it's is_used will be false
func TestChoose(t *testing.T) {
	var url model.ShortURL

	u := store.NewShortURL(db.New(config.Read().Database))
	result := u.Choose()

	u.DB.Raw("SELECT * FROM short_urls WHERE url = ?", result).Scan(&url) //O(lgn)

	if !url.IsUsed {
		t.Errorf("After choosing a short URL it's is_used doesn't change properly")
	}
}

func TestMapInteraction(t *testing.T) {
	m := store.NewMap(db.New(config.Read().Database))
	s := model.Map{
		LongURL:        "LongURLForTest",
		ShortURL:       "ShortURLForTest",
		ExpirationTime: time.Date(2021, 03, 19, 15, 10, 15, 34, time.UTC),
	}

	m.DB.Exec("DELETE FROM maps WHERE long_url = ?", s.LongURL)

	if err := m.Insert(s); err != nil {
		t.Errorf("Insert is not done properly")
	}

	r, e := m.Retrieve(s.ShortURL)
	if e != nil {
		t.Errorf("Cannot retrieve from database")
	}

	if !equal(r, s) {
		t.Errorf("Retrieved row is different from inserted")
	}
}

//Insert same short URL
func TestSameShortURL(t *testing.T) {
	m := store.NewMap(db.New(config.Read().Database))
	s := model.Map{
		LongURL:        "LongURLForTestingSameShortURL",
		ShortURL:       "Same",
		ExpirationTime: time.Date(2021, 03, 19, 15, 10, 15, 34, time.UTC),
	}

	m.DB.Exec("DELETE FROM maps WHERE short_url = ?", s.ShortURL)

	if err := m.Insert(s); err == nil {
		t.Errorf("Cannot insert to database")
	}

	if err := m.Insert(s); err == nil {
		t.Errorf("Same shortURL inserted")
	}
}

func equal(f model.Map, s model.Map) bool {
	if f.LongURL == s.LongURL && f.ShortURL == s.ShortURL && f.ExpirationTime.Unix() == s.ExpirationTime.Unix() {
		return true
	}

	return false
}
