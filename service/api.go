package service

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/elahe-dastan/urlShortener/config"
	"github.com/elahe-dastan/urlShortener/metric"
	"github.com/elahe-dastan/urlShortener/request"
	"github.com/elahe-dastan/urlShortener/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API struct {
	Map      store.Map
	ShortURL store.ShortURL
}

func (a API) Run(cfg config.LogFile) {
	e := echo.New()
	e.POST("/urls", a.Mapping)
	e.GET("/redirect/:shortURL", a.Redirect)

	go func() {
		metric.Monitor()
	}()

	f, _ := os.OpenFile(cfg.Address, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: f}))
	e.Logger.Fatal(e.Start(":8080"))
}

func (a API) Mapping(c echo.Context) error {
	var newMap request.Map

	if err := c.Bind(&newMap); err != nil {
		return err
	}

	if !newMap.Validate() {
		return echo.NewHTTPError(http.StatusBadRequest, "This is not a url at all")
	}

	if newMap.ExpirationTime.Before(time.Now()) {
		var duration time.Duration = 5
		newMap.ExpirationTime = time.Now().Add(duration * time.Minute)
	}

	if newMap.ShortURL == "" {
		newMap = a.randomShortURL(newMap)
	} else if !a.customShortURL(newMap) {
		return echo.NewHTTPError(http.StatusConflict, "This short url exists")
	}

	return c.JSON(http.StatusCreated, newMap)
}

func (a API) randomShortURL(new request.Map) request.Map {
	for {
		u := a.ShortURL.Choose()
		log.Print(u)
		new.ShortURL = u

		if err := a.Map.Insert(new.Model()); err == nil {
			return new
		}
	}
}

func (a API) customShortURL(newMap request.Map) bool {
	if !a.ShortURL.Unique(newMap.ShortURL) {
		return false
	}

	if err := a.Map.Insert(newMap.Model()); err != nil {
		return false
	}

	return true
}

func (a API) Redirect(c echo.Context) error {
	shortURL := c.Param("shortURL")
	if !CheckShortURL(shortURL) {
		return echo.NewHTTPError(http.StatusBadRequest, shortURL)
	}

	mapping, err := a.Map.Retrieve(shortURL)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, shortURL)
	}

	return c.JSON(http.StatusFound, mapping.LongURL)
}

func CheckShortURL(shortURL string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z]+$", shortURL)

	return match
}
