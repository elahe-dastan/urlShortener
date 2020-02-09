package services

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
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
	selectedShortURL := db.ChooseShortURLRowLock()
	newMap.ShortURL = selectedShortURL
	db.InsertToMapping(newMap)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMap)
}

func redirection(w http.ResponseWriter, r *http.Request) {
	shortURL := mux.Vars(r)["shortURL"]
	fmt.Print(shortURL)
	mapping := db.RetriveLongURL(shortURL)
	w.WriteHeader(http.StatusFound)

	json.NewEncoder(w).Encode(mapping)
}

