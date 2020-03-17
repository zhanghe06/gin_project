package main

import (
	"fmt"
)

func testMap() {
	fmt.Println("\ntestMap:")
	var mainMap map[string]string

	mainMap = map[string]string{"1": "a", "2": "b", "3": "c"}

	for key, value := range mainMap {
		fmt.Println(key, ":", value)
	}
}

func testNestMap() {
	fmt.Println("\ntestNestMap:")
	var mainMap = make(map[string]interface{})
	mainMap["1"] = "a"
	mainMap["2"] = "b"
	mainMap["3"] = map[string]string{"1": "a", "2": "b", "3": "c"}

	for key, value := range mainMap {
		fmt.Println(key, ":", value)

		// 断言嵌套结构
		if subMap, ok := value.(map[string]string); ok {
			for k, v := range subMap {
				fmt.Println("\t", k, ":", v)
			}
		}
	}
}

func main() {
	testMap()
	testNestMap()
}
