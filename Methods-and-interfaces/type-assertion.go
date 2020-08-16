package main

import "fmt"

func main()  {
	var ir interface{} = "world"
	s := ir.(string)
	fmt.Println(s)

	s, ok := ir.(string)
	fmt.Println(s,ok)

	//assertion check(ok)
	//innterfaceが保持していない型のため、型assertionのcheckが必要になる
	//okにはbool値を返す必要がある（true,false問わず）
	f, ok := ir.(float64)
	fmt.Println(f,ok)

	//assertion check(NGのためコメントアウト)
	//f = ir.(float64)
	//fmt.Println(f)

}