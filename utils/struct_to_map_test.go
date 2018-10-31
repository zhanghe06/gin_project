package utils

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

type ScoreStruct struct {
	Score	int
}


func TestStruct2Map(t *testing.T) {
	var scoreStruct ScoreStruct
	testValue := 2
	scoreStruct.Score = testValue
	scoreMap := Struct2Map(scoreStruct)
	assert.Equal(t, scoreMap, map[string]interface{}{"Score": testValue})
}
