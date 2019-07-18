package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("nihao", "ni"))
	fmt.Println(strings.Contains("nihao", "nihaoa"))
	fmt.Println(strings.TrimLeft("nihao", "ni"))
	fmt.Println(strings.TrimLeft("ninhao", "ni"))  // 注意 hao
	fmt.Println(strings.TrimPrefix("ninhao", "ni"))  // 注意 nhao
	fmt.Println(strings.Trim("  ni hao a ", " "))  // 注意
	fmt.Println(strings.Count("  nihaoa ", "a"))
	fmt.Println(strings.Index("  nihaoa ", "a"))
	fmt.Println(strings.Join([]string{"ni", "hao"}, "-"))
	fmt.Println(strings.Split("ni-hao", "-"))
	fmt.Println(strings.Repeat("nihao", 3))
	fmt.Println(strings.ToLower("NIHAO"))
	fmt.Println(strings.ToUpper("nihao"))
	fmt.Println(strings.ToUpper("nihao"))
	fmt.Println(len("nihao"))
	fmt.Println("nihao"[3])
}
