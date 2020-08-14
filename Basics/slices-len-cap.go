package main

import (
	"fmt"
)

func main() {
	s := []int{1,2,3,4,5,6}
	fmt.Println(s)

	s = s[:0]
	fmt.Println(s)

	s = s[:4]
	fmt.Println(s)

	s = s[2:]
	fmt.Println(s)

	s = append(s, 10)
	fmt.Println(s)

	v := []int{10,11,12,13,14,15}
	fmt.Println(v)

	sv := append(s,v...)
	fmt.Println(sv)

	// replase
	nav := 5
	sv = append(sv[:nav],sv[nav:]...)
	sv[nav] = 99
	fmt.Println(sv)

	// insert
	xav := 2
	xv := []int{98,99,101,102,103}
	xv = append(xv[:xav+1],xv[xav:]...)
	xv[xav] = 100
	fmt.Println(xv)

}