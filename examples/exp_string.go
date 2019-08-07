package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("nihao", "ni"))          // true
	fmt.Println(strings.Contains("nihao", "nihaoa"))      // false
	fmt.Println(strings.TrimLeft("nihao", "ni"))          // hao
	fmt.Println(strings.TrimLeft("ninhao", "ni"))         // hao
	fmt.Println(strings.TrimPrefix("ninhao", "ni"))       // nhao
	fmt.Println(strings.Trim("  ni hao a ", " "))         // ni hao a
	fmt.Println(strings.Count("  nihaoa ", "a"))          // 2
	fmt.Println(strings.Index("  nihaoa ", "a"))          // 5
	fmt.Println(strings.Join([]string{"ni", "hao"}, "-")) // ni-hao
	fmt.Println(strings.Join([]string{"ni", ""}, "-"))    // ni-
	fmt.Println(strings.Split("ni-hao", "-"))             // [ni hao]
	fmt.Println(strings.Repeat("nihao", 3))               // nihaonihaonihao
	fmt.Println(strings.ToLower("NIHAO"))                 // nihao
	fmt.Println(strings.ToUpper("nihao"))                 // NIHAO
	fmt.Println(len("nihao"))                             // 5
	fmt.Println("nihao"[3])                               // 97
}
