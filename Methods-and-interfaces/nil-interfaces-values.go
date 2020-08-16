package main

import (
	"fmt"
)

type Inter4 interface {
	M01()
}

type Notes struct {
	name string
	age   int
	from  string
}

func (n *Notes) M01() {
	if n == nil {
		fmt.Println("none")
		return
	}
	fmt.Printf("name：%v\n", n.name)
	fmt.Printf("age：%v\n", n.age)
	fmt.Printf("sex：%v\n", n.from)

}

//original
func main() {
	var ic Inter4
	var note *Notes
	ic = note
	describe01(ic)
	ic.M01()

	ic = &Notes{"testuser", 30, "ja"}
	describe01(ic)
	ic.M01()

	ic = &Notes{"testuser2", 10, "En"}
	describe01(ic)
	ic.M01()
}

func describe01(ics Inter4) {
	fmt.Printf("(%v %T)\n", ics, ics)
}
