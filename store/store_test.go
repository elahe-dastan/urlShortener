package store

import (
	"log"
	"testing"
	"time"

	"github.com/elahe-dastan/urlShortener/config"
	"github.com/elahe-dastan/urlShortener/db"
	"github.com/elahe-dastan/urlShortener/model"
	"github.com/stretchr/testify/assert"
)

//Choose a random row and check if it's is_used will be false
func TestChoose(t *testing.T) {
	var url model.ShortURL

	u := NewShortURL(db.New(config.Read().Database))
	result := u.Choose()

	err := u.DB.QueryRow("SELECT * FROM short_url WHERE url = $1", result).Scan(&url.URL, &url.IsUsed)
	if err != nil {
		log.Fatal(err)
	}

	assert.True(t, url.IsUsed, "After choosing a short URL it's is_used doesn't change properly")
}

func TestMapInteraction(t *testing.T) {
	m := NewMap(db.New(config.Read().Database))
	s := model.Map{
		LongURL:        "LongURLForTest",
		ShortURL:       "ShortURLForTest",
		ExpirationTime: time.Date(2021, 03, 19, 15, 10, 15, 34, time.UTC),
	}

	_, err := m.DB.Exec("DELETE FROM map WHERE long_url = $1", s.LongURL)
	if err != nil {
		log.Fatal(err)
	}

	assert.Nil(t, m.Insert(s), "Insert is not done properly")

	r, e := m.Retrieve(s.ShortURL)
	assert.Nil(t, e, "Cannot retrieve from database")

	if !equal(r, s) {
		t.Errorf("Retrieved row is different from inserted")
	}
}

//Insert same short URL
//This should never happen
//func TestSameShortURL(t *testing.T) {
//	m := NewMap(db.New(config.Read().Database))
//	s := model.Map{
//		LongURL:        "LongURLForTestingSameShortURL",
//		ShortURL:       "Same",
//		ExpirationTime: time.Date(2021, 03, 19, 15, 10, 15, 34, time.UTC),
//	}
//
//	_, err := m.DB.Exec("DELETE FROM map WHERE short_url = $1", s.ShortURL)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	assert.Nil(t, m.Insert(s), "Cannot insert to database")
//
//	assert.NotNil(t, m.Insert(s), "Same shortURL inserted")
//}

func equal(f model.Map, s model.Map) bool {
	if f.LongURL == s.LongURL && f.ShortURL == s.ShortURL && f.ExpirationTime.Unix() == s.ExpirationTime.Unix() {
		return true
	}

	return false
}
