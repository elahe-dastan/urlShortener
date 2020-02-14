package services

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
	"urlShortener/KGS"
	"urlShortener/db"
	"urlShortener/models"
)

func RunServices()  {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/urls", MapToShortURL)
	router.HandleFunc("/redirect/{shortURL}", RedirectToLongURL)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func MapToShortURL(w http.ResponseWriter, r *http.Request) {
	var newMap models.ShortToLongURLMap
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "can not read the request due to the following err\n :%s", err)
	}

	json.Unmarshal(reqBody, &newMap)

	if !CheckURL(newMap) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//this part of code doesn't look good
	if newMap.ExpirationTime.Before(time.Now()) {
		newMap.ExpirationTime = time.Now().Add(2*time.Minute)
	}

	if newMap.ShortURL == "" {
		newMap = AssignRandomShortURL(newMap)
	}else {
		if !AssignCustomShortURL(newMap) {
			w.WriteHeader(http.StatusConflict)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMap)
}

func AssignRandomShortURL(newMap models.ShortToLongURLMap) models.ShortToLongURLMap {
	for true {
		selectedShortURL := db.ChooseShortURLTransaction()
		newMap.ShortURL = selectedShortURL
		if err := db.InsertToMapping(newMap);err == nil {
			return newMap
		}
	}
	return newMap
}

func AssignCustomShortURL(newMap models.ShortToLongURLMap) bool {
	if err := db.InsertToMapping(newMap);err != nil {
		return false
	}
	return true
}

func RedirectToLongURL(w http.ResponseWriter, r *http.Request) {
	shortURL := mux.Vars(r)["shortURL"]
	if !CheckShortURL(shortURL) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mapping, err := db.RetrieveLongURL(shortURL)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, mapping.LongURL, http.StatusFound)
	json.NewEncoder(w).Encode(mapping)
}

func CheckURL(newMap models.ShortToLongURLMap) bool {
	_, err := url.ParseRequestURI(newMap.LongURL)
	if err != nil {
		return false
	}
	return true
}

func CheckShortURL(shortURL string) bool {
	//check the length of shortURL
	if len(shortURL) != KGS.LengthOfShortURL {
		return false
	}

	match, _ := regexp.MatchString("^[a-zA-Z]+$", shortURL)
	return match
}

