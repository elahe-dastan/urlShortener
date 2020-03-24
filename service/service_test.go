package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/elahe-dastan/urlShortener/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestInvalidCharacterInShortURL(t *testing.T) {
	shortURL := "34"
	result := CheckShortURL(shortURL)

	assert.False(t, result, "Validtion was incorrect")
}

func TestDBInteraction(t *testing.T) {
	a := API{
		Map:      mock.Map{Maps: map[string]string{}},
		ShortURL: mock.ShortURL{ShortURLs: map[string]bool{"ahs": false}},
	}

	mapJSON := `{"LongURL":"https://www.geeksforgeeks.org",
				 "ShortURL":"ahs",
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

	assert.Nil(t, err, "Cannot read body")

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	fmt.Println(string(body))

	assert.Nil(t, resp.Body.Close(), "Cannot close body")

	ec := echo.New()
	ec.GET("/redirect/:shortURL", a.Redirect)

	requ := httptest.NewRequest(http.MethodGet, "/redirect/ahs", nil)
	reco := httptest.NewRecorder()
	ec.ServeHTTP(reco, requ)

	assert.Equal(t, http.StatusFound, reco.Code)
}
