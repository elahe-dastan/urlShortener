package store

import (
	"errors"

	"github.com/elahe-dastan/urlShortener_KGS/model"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/client_golang/prometheus"
)

type Map struct {
	DB      *gorm.DB
	Counter prometheus.Counter
}

func NewMap(DB *gorm.DB) *Map {
	return &Map{DB: DB,
		Counter: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "shorturl",
			Name:      "counter",
		}),
	}
}

// Creates a table in the database that matches the Map table and puts a trigger on it which deletes the
// rows that have expired after each insert
func (m Map) Create() {
	if m.DB.HasTable(&model.Map{}) {
		return
	}

	m.DB.Debug().AutoMigrate(&model.Map{})

	m.DB.Exec("create or replace function delete_expired_row() " +
		"returns trigger as " +
		"$BODY$ " +
		"begin " +
		"delete from maps where expiration_time < NOW(); " +
		"return null; " +
		"end; " +
		"$BODY$ " +
		"LANGUAGE plpgsql;" +
		"create trigger delete_expired_rows " +
		"after insert " +
		"on maps " +
		"for each row " +
		"execute procedure delete_expired_row();")

	m.DB.Exec("create or replace function give_back_url() " +
		"returns trigger as " +
		"$BODY$ " +
		"begin " +
		"update short_urls set is_used=false where url=old.short_url; " +
		"return null; " +
		"end; " +
		"$BODY$ " +
		"LANGUAGE plpgsql;" +
		"create trigger give_back_url " +
		"after delete " +
		"on maps " +
		"for each row " +
		"execute procedure give_back_url();")
}

// Inserts a Map model in the database
func (m Map) Insert(urlMap model.Map) error {
	err := m.DB.Create(&urlMap).Error
	m.Counter.Inc()
	return err
}

// Gets a short url as parameter and returns a Map model
func (m Map) Retrieve(url string) (model.Map, error) {
	var mapping model.Map

	m.DB.Raw("SELECT * from maps WHERE url = ?;", url).Scan(&mapping) //O(lgn)

	var err error
	if mapping.LongURL == "" {
		err = errors.New("this short URL doesn't exist in the database")
	}

	return mapping, err
}
