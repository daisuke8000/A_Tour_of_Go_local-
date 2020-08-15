package main

import "fmt"

type InterF interface {
	M()
}

type Typo struct {
	S string
}

func (t Typo) M() {
	fmt.Println(t.S)
}

func main() {
	var inter InterF = Typo{"Hello"}
	fmt.Println(inter)
	inter.M()
}
