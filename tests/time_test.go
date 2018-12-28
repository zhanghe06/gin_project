package tests

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"time"
)

func getCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func getCurrentUTCTime() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}

func TestGetCurrentTime(t *testing.T) {
	currentTime := getCurrentTime()
	assert.Equal(t, currentTime, time.Now().Format("2006-01-02 15:04:05"))
}

func TestGetCurrentUTCTime(t *testing.T) {
	currentUTCTime := getCurrentUTCTime()
	assert.Equal(t, currentUTCTime, time.Now().UTC().Format("2006-01-02 15:04:05"))
}
