package main

import "fmt"

func main() {
	var z []int
	a := make([]int, 5)    // len(a)=5
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	c := b[0:5]
	fmt.Println(z, len(z), cap(z), z == nil) // [] 0 0 true
	fmt.Println(a, len(a), cap(a), a == nil) // [0 0 0 0 0] 5 5 false
	fmt.Println(b, len(b), cap(b), b == nil) // [] 0 5 false
	fmt.Println(c, len(c), cap(c), c == nil) // [0 0 0 0 0] 5 5 false
}
