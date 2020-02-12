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
	"urlShortener/KGS"
	"urlShortener/db"
	"urlShortener/models"
)

func RunServices()  {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/urls", mapLongURL)
	router.HandleFunc("/redirect/{shortURL}", redirection)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func mapLongURL(w http.ResponseWriter, r *http.Request) {
	var newMap models.ShortURLMap
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the long URL")
	}

	json.Unmarshal(reqBody, &newMap)
	if !CheckURL(newMap) {
		w.WriteHeader(http.StatusBadRequest)
		//json.NewEncoder(w).Encode(newMap)
	}

	selectedShortURL := db.ChooseShortURLTransaction()
	newMap.ShortURL = selectedShortURL
	db.InsertToMapping(newMap)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMap)
}

func redirection(w http.ResponseWriter, r *http.Request) {
	shortURL := mux.Vars(r)["shortURL"]
	if !CheckShortURL(shortURL) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	mapping := db.RetriveLongURL(shortURL)
	if mapping.LongURL == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	http.Redirect(w, r, mapping.LongURL, http.StatusFound)
	json.NewEncoder(w).Encode(mapping)
}

func CheckURL(newMap models.ShortURLMap) bool {
	//check url length
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

