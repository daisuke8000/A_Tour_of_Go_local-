package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

//original
func bx(fnc func(string, string) string) string {
	return fnc("Thanks", "You!")
}

func main() {
	// 三角形の斜辺の長さの計算式
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(2, 1))
	fmt.Println(compute(hypot))

	//original
	ax := func(s1, s2 string) string {
		return (s1 + " " + s2)
	}
	fmt.Println(ax("Hello", "World!"))
	fmt.Println(bx(ax))
}