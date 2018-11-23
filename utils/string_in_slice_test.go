package utils

import (
	"github.com/magiconair/properties/assert"
	"testing"
)


func TestStringInSlice(t *testing.T) {
	assert.Equal(t, StringInSlice("a", []string{"a", "b", "c"}), true)
	assert.Equal(t, StringInSlice("d", []string{"a", "b", "c"}), false)
}
