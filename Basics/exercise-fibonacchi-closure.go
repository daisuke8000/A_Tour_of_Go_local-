package main

import "fmt"

func fibonacci() func() int {
	i0, i1 := -1, 1
	return func() int {
		i0, i1 = i1, (i0 + i1)
		return i1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
