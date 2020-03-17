package main

import (
	"fmt"
	"strconv"
	"time"
)


func mapDemoBuiltIn() {
	// n太小时不会（比如20以内），因机器而异
	// fatal error: concurrent map read and map write
	n := 20
	m := make(map[string]int)
	//forever := make(chan bool)

	go func() {
		for i := 0; i < n; i++ {
			m[strconv.Itoa(i)] = i // write
		}
	}()

	go func() {
		for i := 0; i < n; i++ {
			fmt.Println(i, m[strconv.Itoa(i)]) // read
		}
	}()

	time.Sleep(time.Second * 5)
	//<-forever
}

func main() {
	mapDemoBuiltIn()
}
