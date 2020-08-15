package main

import (
	"fmt"
	"math"
)

type Vertexs struct {
	X, Y float64
}

func (v Vertexs) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertexs) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertexs{3, 5}
	v.Scale(10)
	fmt.Println(v.Abs())
}
