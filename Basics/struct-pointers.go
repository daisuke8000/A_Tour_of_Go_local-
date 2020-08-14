package main

import "fmt"

type Anyone struct {
	X int
	Y int
}

func main() {
	v := Anyone{3,5}
	d := &v
	fmt.Println(d.Y)
	d.Y = 1e9
	fmt.Println(v)
}