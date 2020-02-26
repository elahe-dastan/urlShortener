package tests

import (
	"testing"
	"urlShortener/models"
	"urlShortener/services"
)

func TestUniCode(t *testing.T)  {
	newMap := models.ShortToLongURLMap{LongURL:"https://fa.wikipedia.org/wiki/%D8%AA%D9%87%D8%B1%D8%A7%D9%86"}
	result := services.CheckURL(newMap)
	if result == false {
		t.Errorf("Validtion was incorrect")
	}
}

func TestEmptyURL(t *testing.T) {
	newMap := models.ShortToLongURLMap{LongURL:""}
	result := services.CheckURL(newMap)
	if result == true {
		t.Errorf("Validtion was incorrect")
	}
}

func TestInvalidURL(t *testing.T) {
	newMap := models.ShortToLongURLMap{LongURL:"sdf"}
	result := services.CheckURL(newMap)
	if result == true {
		t.Errorf("Validtion was incorrect")
	}
}