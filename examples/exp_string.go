package main

import (
	"fmt"
	"strings"
)

func testString01() {
	var str1 string // 默认值为空字符串 ""
	str1 = `hello world`
	str2 := "你好世界"

	str := str1 + " " + str2 // 字符串连接
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str) // 输出：hello world 你好世界

	// 遍历字符串
	l := len(str)
	fmt.Println(l)
	for i := 0; i < l; i++ {
		chr := str[i]
		fmt.Println(i, chr) // 输出字符对应的编码数字
	}
}

func testString02() {
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
	fmt.Println(strings.Replace("ni-hao-", "-", "", -1))  // nihao
	fmt.Println(strings.ReplaceAll("ni-hao-", "-", ""))   // nihao
	fmt.Println(strings.Repeat("nihao", 3))               // nihaonihaonihao
	fmt.Println(strings.ToLower("NIHAO"))                 // nihao
	fmt.Println(strings.ToUpper("nihao"))                 // NIHAO
	fmt.Println(len("nihao"))                             // 5
	fmt.Println("nihao"[3])                               // 97
}

func main() {
	testString01()
	testString02()
}
