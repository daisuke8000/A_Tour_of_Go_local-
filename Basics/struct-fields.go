package main

import "fmt"

type Vertexs struct {
	X int
	Y int
}

func main() {
	v := Vertexs{1,2}
	v.X = 4
	fmt.Println(v.X)
}
