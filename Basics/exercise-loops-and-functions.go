package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(3)
	z -= (z*z - x) / (2*z)
	return z
}

func main() {
	fmt.Println(Sqrt(9))
}
