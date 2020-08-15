package main

import (
	"fmt"
	"math"
)

type Inter1 interface {
	M()
}

type Typo1 struct {
	S string
}

type Fob float64

func (t *Typo1) M() {
	fmt.Println(t.S)
}

func (f Fob) M() {
	fmt.Println(f)
}

func main() {
	var ifc Inter1
	ifc = &Typo1{"True"}
	describe(ifc)
	ifc.M()

	ifc = Fob(math.Pi)
	describe(ifc)
	ifc.M()
}

func describe(i Inter1) {
	fmt.Printf("(%v,%T)\n", i, i)
}
