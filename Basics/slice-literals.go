package main

import "fmt"

type Core struct {
	i string
	v int
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(a)

	b := []bool{true, false, true, true, false, true}
	fmt.Println(b)

	c := Core{"hoge", 1}
	fmt.Println(c)

	d := []struct {
		i string
		b bool
	}{
		{"fuga", true},
		{"hogege", false},
	}
	fmt.Println(d[0])
}
