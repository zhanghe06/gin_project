package main

import "fmt"

// 错误写法
// examples/exp_var.go:11:14: undefined: a
//func main()  {
//	if true {
//		a := 1
//	} else {
//		a := 2
//	}
//	fmt.Println(a)
//}

func main()  {
	var a int32
	if true {
		a = 1
	} else {
		a = 2
	}
	fmt.Println(a)
}
