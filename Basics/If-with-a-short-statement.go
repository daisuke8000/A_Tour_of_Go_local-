package main

import (
	"fmt"
	"math"
)

func pows(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v > lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(pows(5, 2, 10))
	fmt.Println(pows(3,3, 28))
}
