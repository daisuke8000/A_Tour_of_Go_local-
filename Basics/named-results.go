package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 5
	y = sum - 10 * 8
	return
}

func main() {
	va, vb := split(22)
	fmt.Println(va, vb)
	fmt.Println(split(17))
}
