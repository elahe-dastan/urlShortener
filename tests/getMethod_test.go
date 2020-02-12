package tests

import (
	"testing"
	"urlShortener/services"
)

func TestInvalidCharacterInShortURL(t *testing.T)  {
	shortURL := "34"
	result := services.CheckShortURL(shortURL)
	if result == true {
		t.Errorf("Validtion was incorrect")
	}
}

func TestInvalidShortURLLength(t *testing.T)  {
	shortURL := "fsg"
	result := services.CheckShortURL(shortURL)
	if result == true {
		t.Errorf("Validtion was incorrect")
	}
}
