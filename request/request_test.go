package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniCode(t *testing.T) {
	newMap := Map{LongURL: "https://fa.wikipedia.org/wiki/%D8%AA%D9%87%D8%B1%D8%A7%D9%86"}
	assert.True(t, newMap.Validate(), "Validtion was incorrect")
}

func TestEmptyURL(t *testing.T) {
	newMap := Map{LongURL: ""}
	assert.False(t, newMap.Validate(), "Validtion was incorrect")
}

func TestInvalidURL(t *testing.T) {
	newMap := Map{LongURL: "sdf"}
	assert.False(t, newMap.Validate(), "Validtion was incorrect")
}
