package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"
	"urlShortener/generator"
	"urlShortener/db"
	"urlShortener/model"
)

func Run()  {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/urls", mapping)
	router.HandleFunc("/redirect/{shortURL}", redirect)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func mapping(w http.ResponseWriter, r *http.Request) {
	var new model.Map
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "can not read the request due to the following err\n :%s", err)
	}

	json.Unmarshal(reqBody, &new)

	if !CheckLongURL(new) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// this part of code doesn't look good
	if new.ExpirationTime.Before(time.Now()) {
		new.ExpirationTime = time.Now().Add(2*time.Minute)
	}

	if new.ShortURL == "" {
		new = randomShortURL(new)
	}else {
		if !customShortURL(new) {
			w.WriteHeader(http.StatusConflict)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(new)
}

func randomShortURL(new model.Map) model.Map {
	for true {
		url := db.ChooseShortURL()
		new.ShortURL = url
		if err := db.InsertMap(new);err == nil {
			return new
		}
	}
	return new
}

func customShortURL(newMap model.Map) bool {
	if err := db.InsertMap(newMap);err != nil {
		return false
	}
	return true
}

func redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := mux.Vars(r)["shortURL"]
	if !CheckShortURL(shortURL) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mapping, err := db.Retrieve(shortURL)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, mapping.LongURL, http.StatusFound)
	json.NewEncoder(w).Encode(mapping)
}

func CheckLongURL(newMap model.Map) bool {
	_, err := url.ParseRequestURI(newMap.LongURL)
	if err != nil {
		return false
	}
	return true
}

func CheckShortURL(shortURL string) bool {
	//check the length of shortURL
	if len(shortURL) != generator.LengthOfShortURL {
		return false
	}

	match, _ := regexp.MatchString("^[a-zA-Z]+$", shortURL)
	return match
}

