package main

import "fmt"

func main()  {
	pochi := make([]int,10)
	for i := range pochi {
		pochi[i] = 2 << uint(i)
		fmt.Println(pochi[i])
	}
	for _, value := range pochi {
		fmt.Printf("%d\n", value)
	}
}