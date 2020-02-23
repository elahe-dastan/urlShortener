package test

import (
	"testing"
	"urlShortener/service"
)

func TestInvalidCharacterInShortURL(t *testing.T)  {
	shortURL := "34"
	result := service.CheckShortURL(shortURL)
	if result == true {
		t.Errorf("Validtion was incorrect")
	}
}

func TestInvalidShortURLLength(t *testing.T)  {
	shortURL := "fsg"
	result := service.CheckShortURL(shortURL)
	if result == true {
		t.Errorf("Validtion was incorrect")
	}
}
