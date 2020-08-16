package main

import "fmt"

type InterF3 interface {
	Mon()
}

type T2 struct {
	S2 string
}

func (t2 *T2) Mon() {
	if t2 == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t2.S2)
}

func main() {
	var hell InterF3
	var t2 *T2
	hell = t2
	describe5(hell)
	hell.Mon()

	hell = &T2{"HELLO"}
	describe5(hell)
	hell.Mon()
}

func describe5(sen InterF3) {
	fmt.Printf("(%v %T)\n", sen, sen)
}
