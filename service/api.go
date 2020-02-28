package service

import (
	"encoding/json"
	"fmt"
	"github.com/elahe-dastan/urlShortener_KGS/db"
	"github.com/elahe-dastan/urlShortener_KGS/generator"
	"github.com/elahe-dastan/urlShortener_KGS/middleware"
	"github.com/elahe-dastan/urlShortener_KGS/model"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

func Run()  {
	//router := mux.NewRouter().StrictSlash(true)
	mux := &http.ServeMux{}
	mux.HandleFunc("/urls", mapping)
	mux.HandleFunc("/redirect/{shortURL}", redirect)
	handler := middleware.LogRequestHandler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func mapping(w http.ResponseWriter, r *http.Request) {
	var newMap model.Map
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Header().Add("err",err.Error())
		fmt.Fprintf(w, "can not read the request due to the following err\n :%s", err)
	}

	json.Unmarshal(reqBody, &newMap)

	if !CheckLongURL(newMap) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// this part of code doesn't look good
	if newMap.ExpirationTime.Before(time.Now()) {
		newMap.ExpirationTime = time.Now().Add(2*time.Minute)
	}

	if newMap.ShortURL == "" {
		newMap = randomShortURL(newMap)
	}else {
		if !customShortURL(newMap) {
			w.WriteHeader(http.StatusConflict)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMap)
}

func randomShortURL(new model.Map) model.Map {
	for  {
		u := db.ChooseShortURL()
		new.ShortURL = u
		if err := db.InsertMap(new);err == nil {
			return new
		}
	}
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
	return err != nil
}

func CheckShortURL(shortURL string) bool {
	//check the length of shortURL
	if len(shortURL) != generator.LengthOfShortURL {
		return false
	}

	match, _ := regexp.MatchString("^[a-zA-Z]+$", shortURL)
	return match
}

