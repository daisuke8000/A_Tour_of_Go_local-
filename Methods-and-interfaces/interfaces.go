package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := Myfloat(-math.Sqrt2)
	v := Vertexs0{3, 4}
	a = f
	a = &v
	fmt.Println(a)
	//a = v
	fmt.Println(a.Abs())

	// interfaceの使い方１(なんでもはいる)
	var i interface{}
	i = 100
	fmt.Println(i)
	i = "All-Ok"
	fmt.Println(i)
	// 取り出し方が特殊
	str, tf := i.(string)
	fmt.Println(str,tf)

	// interfaceの使い方２（関数まとめたもの）

}

type Myfloat float64

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertexs0 struct {
	X, Y float64
}

func (v *Vertexs0) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
