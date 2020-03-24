package mock

import (
	"github.com/elahe-dastan/urlShortener/model"
)

type Map struct {
	Maps map[string]string
}

type ShortURL struct {
	ShortURLs map[string]bool
}

func (m Map) Insert(urlMap model.Map) error {
	m.Maps[urlMap.ShortURL] = urlMap.LongURL

	return nil
}

// Gets a short url as parameter and returns a Map model
func (m Map) Retrieve(url string) (model.Map, error) {
	lu := m.Maps[url]

	mapping := model.Map{
		LongURL:  lu,
		ShortURL: url,
	}

	return mapping, nil
}

func (s ShortURL) Unique(shortURL string) bool {
	return !s.ShortURLs[shortURL]
}

func (s ShortURL) Save() {

}

func (s ShortURL) Choose() string {
	return "something"
}
