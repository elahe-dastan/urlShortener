package test

import (
	"testing"
	"urlShortener/model"
	"urlShortener/service"
)

func TestUniCode(t *testing.T)  {
	newMap := model.Map{LongURL: "https://fa.wikipedia.org/wiki/%D8%AA%D9%87%D8%B1%D8%A7%D9%86"}
	result := service.CheckLongURL(newMap)
	if result == false {
		t.Errorf("Validtion was incorrect")
	}
}

func TestEmptyURL(t *testing.T) {
	newMap := model.Map{LongURL: ""}
	result := service.CheckLongURL(newMap)
	if result == true {
		t.Errorf("Validtion was incorrect")
	}
}

func TestInvalidURL(t *testing.T) {
	newMap := model.Map{LongURL: "sdf"}
	result := service.CheckLongURL(newMap)
	if result == true {
		t.Errorf("Validtion was incorrect")
	}
}