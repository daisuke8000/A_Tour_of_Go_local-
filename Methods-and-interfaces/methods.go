package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (p Vertex) Abs(i string) (string,float64) {
	return i, math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	p := Vertex{3, 4}
	fmt.Println(p.Abs("絶対値："))
}
