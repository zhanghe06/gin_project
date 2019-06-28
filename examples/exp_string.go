package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("nihao", "ni"))
	fmt.Println(strings.Contains("nihao", "nihaoa"))
	fmt.Println(strings.TrimLeft("nihao", "ni"))
	fmt.Println(strings.TrimLeft("ninhao", "ni"))  // 注意
	fmt.Println(strings.TrimPrefix("ninhao", "ni"))  // 注意
	fmt.Println(strings.Trim("  ni hao a ", " "))  // 注意
}
