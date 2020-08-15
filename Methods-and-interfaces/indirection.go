package main

import "fmt"

type Gore struct {
	X,Y float64
}

func (v *Gore) Scale(f float64)  {
	v.X = v.X + f
	v.Y = v.Y * f
}

func Verocica(x *Gore, f float64){
	x.X = f/2
	x.Y = f/3
}

func main()  {
	v := Gore{1,2}
	v.Scale(15)
	Verocica(&v,17)

	p := Gore{3,4}
	p.Scale(8)
	Verocica(&p,21)

	fmt.Println(v,p)
}