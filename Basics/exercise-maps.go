package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := map[string]int{}
	for _, v := range strings.Fields(s) {
		result[v]++
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
