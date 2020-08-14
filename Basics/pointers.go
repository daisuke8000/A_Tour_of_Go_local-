package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	q := &j
	fmt.Println(*q)
	*q = 98
	fmt.Println(j)
}
