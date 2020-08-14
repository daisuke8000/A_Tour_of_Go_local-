package main

import "fmt"

func main() {
	var s []int
	PrintSlice(s)

	s = append(s, 0)
	PrintSlice(s)

	s = append(s, 1)
	PrintSlice(s)

	s = append(s, 2, 3, 4)
	PrintSlice(s)
}

func PrintSlice(x []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
}
