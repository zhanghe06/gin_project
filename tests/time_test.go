package tests

import (
	"fmt"
	"testing"
	"time"
)

func getCurrentTime() string {
	return time.Now().Format("2006-01-06 15:04:05")
}

func getCurrentUTCTime() string {
	return time.Now().UTC().Format("2006-01-06 15:04:05")
}

func TestTime(t *testing.T) {
	fmt.Println(getCurrentTime())
	fmt.Println(getCurrentUTCTime())
}
