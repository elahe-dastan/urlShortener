package mock

import (
	"github.com/elahe-dastan/urlShortener/model"
)

type MockMap struct {
	Maps map[string]string
}

func (m MockMap) Insert(urlMap model.Map) error {
	m.Maps[urlMap.ShortURL] = urlMap.LongURL

	return nil
}

// Gets a short url as parameter and returns a Map model
func (m MockMap) Retrieve(url string) (model.Map, error) {
	lu := m.Maps[url]

	mapping := model.Map{
		LongURL:  lu,
		ShortURL: url,
	}

	return mapping, nil
}
