package store

import (
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

	u.DB.QueryRow("SELECT * FROM short_urls WHERE url = $1", result).Scan(&url) //O(lgn)

	assert.True(t, url.IsUsed, "After choosing a short URL it's is_used doesn't change properly")
}

func TestMapInteraction(t *testing.T) {
	m := NewMap(db.New(config.Read().Database))
	s := model.Map{
		LongURL:        "LongURLForTest",
		ShortURL:       "ShortURLForTest",
		ExpirationTime: time.Date(2021, 03, 19, 15, 10, 15, 34, time.UTC),
	}

	m.DB.Exec("DELETE FROM maps WHERE long_url = $1", s.LongURL)

	assert.Nil(t, m.Insert(s), "Insert is not done properly")

	r, e := m.Retrieve(s.ShortURL)
	assert.Nil(t, e, "Cannot retrieve from database")

	if !equal(r, s) {
		t.Errorf("Retrieved row is different from inserted")
	}
}

//Insert same short URL
func TestSameShortURL(t *testing.T) {
	m := NewMap(db.New(config.Read().Database))
	s := model.Map{
		LongURL:        "LongURLForTestingSameShortURL",
		ShortURL:       "Same",
		ExpirationTime: time.Date(2021, 03, 19, 15, 10, 15, 34, time.UTC),
	}

	m.DB.Exec("DELETE FROM maps WHERE short_url = $1", s.ShortURL)

	assert.Nil(t, m.Insert(s), "Cannot insert to database")

	assert.NotNil(t, m.Insert(s), "Same shortURL inserted")
}

func equal(f model.Map, s model.Map) bool {
	if f.LongURL == s.LongURL && f.ShortURL == s.ShortURL && f.ExpirationTime.Unix() == s.ExpirationTime.Unix() {
		return true
	}

	return false
}
