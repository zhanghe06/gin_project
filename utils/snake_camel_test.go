package utils

import (
	"github.com/magiconair/properties/assert"
	"testing"
)


func TestSnakeString(t *testing.T) {
	assert.Equal(t, "snake_string", SnakeString("SnakeString"))
}

func TestCamelString(t *testing.T) {
	assert.Equal(t, "SnakeString", CamelString("snake_string"))
}
