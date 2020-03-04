package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/elahe-dastan/urlShortener_KGS/config"
	"github.com/elahe-dastan/urlShortener_KGS/generator"
	"github.com/elahe-dastan/urlShortener_KGS/middleware"
	"github.com/elahe-dastan/urlShortener_KGS/model"
	"github.com/elahe-dastan/urlShortener_KGS/store"
	"github.com/gorilla/mux"
)

type API struct {
	Map      store.Map
	ShortURL store.ShortURL
}

func (a API) Run(constant config.LogFile) {
	//router := mux.NewRouter().StrictSlash(true)
	m := &http.ServeMux{}
	m.HandleFunc("/urls", a.mapping)
	m.HandleFunc("/redirect/{shortURL}", a.redirect)

	c := middleware.Configuration{Config: constant}
	handler := c.LogRequestHandler(m)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func (a API) mapping(w http.ResponseWriter, r *http.Request) {
	var newMap model.Map

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Add("err", err.Error())
		fmt.Print(w, "can not read the request due to the following err\n :%s", err)
	}

	err = json.Unmarshal(reqBody, &newMap)
	if err != nil {
		log.Fatal(err)
	}

	if !CheckLongURL(newMap) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// this part of code doesn't look good
	if newMap.ExpirationTime.Before(time.Now()) {
		var duration time.Duration
		duration = 5
		newMap.ExpirationTime = time.Now().Add(duration * time.Minute)
	}

	if newMap.ShortURL == "" {
		newMap = a.randomShortURL(newMap)
	} else if !a.customShortURL(newMap) {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(newMap); err != nil {
		log.Fatal(err)
	}
}

func (a API) randomShortURL(new model.Map) model.Map {
	for {
		u := a.ShortURL.Choose()
		new.ShortURL = u

		if err := a.Map.Insert(new); err == nil {
			return new
		}
	}
}

func (a API) customShortURL(newMap model.Map) bool {
	if err := a.Map.Insert(newMap); err != nil {
		return false
	}

	return true
}

func (a API) redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := mux.Vars(r)["shortURL"]
	if !CheckShortURL(shortURL) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mapping, err := a.Map.Retrieve(shortURL)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, mapping.LongURL, http.StatusFound)

	if err = json.NewEncoder(w).Encode(mapping); err != nil {
		log.Fatal(err)
	}
}

func CheckLongURL(newMap model.Map) bool {
	_, err := url.ParseRequestURI(newMap.LongURL)

	return err == nil
}

func CheckShortURL(shortURL string) bool {
	//check the length of shortURL
	if len(shortURL) != generator.LengthOfShortURL {
		return false
	}

	match, _ := regexp.MatchString("^[a-zA-Z]+$", shortURL)

	return match
}
