package main

import "fmt"

type Vertex4 struct {
	Lat, Long float64
}

var md = map[string]Vertex4{
	"test1": Vertex4{-23.587676,78.9372},
	"test2": Vertex4{35.96522, 12.763352},
}

func main()  {
	fmt.Println(md)
	gt := map[int]int{0:10}
	fmt.Println(gt)
	gs := make(map[string]int)
	gs["row1"] = 10
	fmt.Println(gs)
}