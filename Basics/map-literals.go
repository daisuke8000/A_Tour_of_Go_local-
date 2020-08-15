package main

import "fmt"

type Vertex3 struct {
	Lat, Long float64
}

var mx = map[string]Vertex3{
	"Bell Labs": Vertex3{
		40.68433, -74.39967,
	},
	"Google": Vertex3{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(mx)
}
