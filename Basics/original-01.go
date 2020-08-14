package main

import (
	"fmt"
)

func main() {
	var i int = 10
	var s = make([]int,0)
	fmt.Println(s)
	for x := 0; x < i; x++ {
		fmt.Println("追加")
		s = append(s, x)
	}
	fmt.Println(s)
}
