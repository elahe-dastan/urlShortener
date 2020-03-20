package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/elahe-dastan/urlShortener/config"
	"github.com/elahe-dastan/urlShortener/db"
	"github.com/elahe-dastan/urlShortener/mocks"
	"github.com/elahe-dastan/urlShortener/service"
	"github.com/elahe-dastan/urlShortener/store"
	"github.com/labstack/echo/v4"
	"github.com/magiconair/properties/assert"
)

func TestInvalidCharacterInShortURL(t *testing.T) {
	shortURL := "34"
	result := service.CheckShortURL(shortURL)

	if result == true {
		t.Errorf("Validtion was incorrect")
	}
}

func TestInvalidShortURLLength(t *testing.T) {
	shortURL := "fsg"
	result := service.CheckShortURL(shortURL)

	if result == true {
		t.Errorf("Validtion was incorrect")
	}
}

func TestMapping(t *testing.T) {
	a := service.API{
		Map:      mocks.MockMap{Maps: map[string]string{}},
		ShortURL: store.NewShortURL(db.New(config.Read().Database)),
	}

	mapJSON := `{"LongURL":"https://www.geeksforgeeks.org",
				 "ShortURL":"aacbaaCp",
				 "ExpirationTime":"2020-06-20T15:09:00.097213128+03:30"}
`
	e := echo.New()
	e.POST("/urls", a.Mapping)

	req := httptest.NewRequest(http.MethodPost, "/urls", strings.NewReader(mapJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	resp := rec.Result()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Cannot read body")
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	fmt.Println(string(body))

	if err := resp.Body.Close(); err != nil {
		t.Errorf("Cannot close body")
	}
}

func TestRedirect(t *testing.T) {
	a := service.API{
		Map:      store.NewMap(db.New(config.Read().Database)),
		ShortURL: store.NewShortURL(db.New(config.Read().Database)),
	}

	e := echo.New()
	e.GET("/redirect/:shortURL", a.Redirect)

	req := httptest.NewRequest(http.MethodGet, "/redirect/aaaaaaCp", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusFound, rec.Code)
}
