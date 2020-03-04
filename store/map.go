package store

import (
	"errors"

	"github.com/elahe-dastan/urlShortener_KGS/model"
	"github.com/jinzhu/gorm"
)

type Map struct {
	DB *gorm.DB
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
}

// Inserts a Map model in the database
func (m Map) Insert(urlMap model.Map) error {
	err := m.DB.Create(&urlMap).Error

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
