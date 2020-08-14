package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on")
	//runtime.GOOS => OSを判定する
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}