package store

import (
	"database/sql"
	"errors"
	"log"

	"github.com/elahe-dastan/urlShortener/metric"
	"github.com/elahe-dastan/urlShortener/model"
	"github.com/prometheus/client_golang/prometheus"
)

type Mapping interface {
	Insert(urlMap model.Map) error
	Retrieve(url string) (model.Map, error)
}

type Map struct {
	DB      *sql.DB
	Counter prometheus.Counter
}

func NewMap(d *sql.DB) Map {
	return Map{DB: d,
		Counter: metric.NewCounter("number_of_requests"),
	}
}

// Creates a table in the database that matches the Map table and puts a trigger on it which deletes the
// rows that have expired after each insert
func (m Map) Create() {
	_, err := m.DB.Exec("CREATE TABLE IF NOT EXISTS map (" +
		"id serial PRIMARY KEY," +
		"long_url VARCHAR NOT NULL," +
		"short_url VARCHAR NOT NULL," +
		"expiration_time TIMESTAMP NOT NULL" +
		");")
	if err != nil {
		log.Fatal("Cannot create map table due to the following error", err.Error())
	}

	m.DB.Exec("create or replace function delete_expired_row() " +
		"returns trigger as " +
		"$BODY$ " +
		"begin " +
		"delete from map where expiration_time < NOW(); " +
		"return null; " +
		"end; " +
		"$BODY$ " +
		"LANGUAGE plpgsql;" +
		"create trigger delete_expired_rows " +
		"after insert " +
		"on map " +
		"for each row " +
		"execute procedure delete_expired_row();")

	m.DB.Exec("create or replace function give_back_url() " +
		"returns trigger as " +
		"$BODY$ " +
		"begin " +
		"update short_url set is_used=false where url=old.short_url; " +
		"return null; " +
		"end; " +
		"$BODY$ " +
		"LANGUAGE plpgsql;" +
		"create trigger give_back_url " +
		"after delete " +
		"on map " +
		"for each row " +
		"execute procedure give_back_url();")
}

// Inserts a Map model in the database
func (m Map) Insert(urlMap model.Map) error {
	_, err := m.DB.Exec("INSERT INTO map (long_url, short_url, expiration_time) VALUES ($1, $2, $3)", urlMap.LongURL, urlMap.ShortURL, urlMap.ExpirationTime)
	m.Counter.Inc()

	return err
}

// Gets a short url as parameter and returns a Map model
func (m Map) Retrieve(url string) (model.Map, error) {
	var mapping model.Map

	m.DB.QueryRow("SELECT * from map WHERE short_url = $1;", url).Scan(&mapping.ID, &mapping.LongURL, &mapping.ShortURL, &mapping.ExpirationTime) //O(lgn)

	var err error
	if mapping.LongURL == "" {
		err = errors.New("this short URL doesn't exist in the database")
	}

	return mapping, err
}
