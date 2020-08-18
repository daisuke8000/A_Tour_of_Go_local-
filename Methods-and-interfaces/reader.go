package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello!!,True_or_false!!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Println(n, err)
		fmt.Printf("n = %v err = %v b = %v\n", n,err,b)
		fmt.Printf("b[:] = %q\n", b[:n])
		if err == io.EOF{
			break
		}
	}
}
