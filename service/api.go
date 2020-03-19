package service

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/elahe-dastan/urlShortener/config"
	"github.com/elahe-dastan/urlShortener/generator"
	"github.com/elahe-dastan/urlShortener/request"
	"github.com/elahe-dastan/urlShortener/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type API struct {
	Map      store.Map
	ShortURL store.ShortURL
}

func (a API) Run(cfg config.LogFile) {
	e := echo.New()
	e.POST("/urls", a.mapping)
	e.GET("/redirect/:shortURL", a.redirect)

	p := &http.ServeMux{}
	p.Handle("/metrics", promhttp.Handler())

	go func() {
		log.Fatal(http.ListenAndServe(":8081", p))
	}()

	f, _ := os.OpenFile(cfg.Address, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: f}))
	e.Logger.Fatal(e.Start(":8080"))
}

func (a API) mapping(c echo.Context) error {
	var newMap request.Map

	if err := c.Bind(&newMap); err != nil {
		return err
	}

	if !newMap.Validate() {
		return c.String(http.StatusBadRequest, "This is not a url at all")
	}

	// this part of code doesn't look good
	if newMap.ExpirationTime.Before(time.Now()) {
		var duration time.Duration = 5
		newMap.ExpirationTime = time.Now().Add(duration * time.Minute)
	}

	if newMap.ShortURL == "" {
		newMap = a.randomShortURL(newMap)
	} else if !a.customShortURL(newMap) {
		return c.String(http.StatusConflict, "This short url exists")
	}

	//w.WriteHeader(http.StatusCreated)
	//
	//if err = json.NewEncoder(w).Encode(newMap); err != nil {
	//	log.Fatal(err)
	//}
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
	if err := a.Map.Insert(newMap.Model()); err != nil {
		return false
	}

	return true
}

func (a API) redirect(c echo.Context) error {
	shortURL := c.Param("shortURL")
	if !CheckShortURL(shortURL) {
		return c.String(http.StatusBadRequest, shortURL)
	}

	mapping, err := a.Map.Retrieve(shortURL)

	if err != nil {
		return c.String(http.StatusNotFound, shortURL)
	}

	return c.String(http.StatusFound, mapping.LongURL)
}

func CheckShortURL(shortURL string) bool {
	//check the length of shortURL
	if len(shortURL) != generator.LengthOfShortURL {
		return false
	}

	match, _ := regexp.MatchString("^[a-zA-Z]+$", shortURL)

	return match
}
