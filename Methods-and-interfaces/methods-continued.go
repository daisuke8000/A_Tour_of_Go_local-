package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (mf MyFloat) Abs() float64 {
	if mf < 0 {
		return float64(-mf)
	}
	return float64(mf)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
