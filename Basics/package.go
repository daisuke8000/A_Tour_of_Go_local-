package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//乱数をSeedで生成
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favolite number is", rand.Intn(10))
}
