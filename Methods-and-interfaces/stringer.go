package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string{
	fmt.Println("test1")
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"test", 19}
	m := Person{"test2", 43}
	fmt.Println(a, m)
}
