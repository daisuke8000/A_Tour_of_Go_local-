package main

import "fmt"

func main()  {
	var vt interface{}
	desc(vt)

	vt = 49
	desc(vt)

	vt = "hogehoge"
	desc(vt)
}

func desc(vts interface{})  {
	fmt.Printf("(%v,%T)\n",vts,vts)
}