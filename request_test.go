package main

import (
	"testing"

	"github.com/elahe-dastan/urlShortener/request"
)

func TestUniCode(t *testing.T) {
	newMap := request.Map{LongURL: "https://fa.wikipedia.org/wiki/%D8%AA%D9%87%D8%B1%D8%A7%D9%86"}
	result := newMap.Validate()

	if result == false {
		t.Errorf("Validtion was incorrect")
	}
}

func TestEmptyURL(t *testing.T) {
	newMap := request.Map{LongURL: ""}
	result := newMap.Validate()

	if result == true {
		t.Errorf("Validtion was incorrect")
	}
}

func TestInvalidURL(t *testing.T) {
	newMap := request.Map{LongURL: "sdf"}
	result := newMap.Validate()

	if result == true {
		t.Errorf("Validtion was incorrect")
	}
}
