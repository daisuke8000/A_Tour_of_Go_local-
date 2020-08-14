package main

import "fmt"

type Vertex1 struct {
	X int
	Y int
}

var (
	v1 = Vertex1{1, 2}
	v2 = Vertex1{X: 5}
	v3 = Vertex1{}
	p  = &Vertex1{1, 2}
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
