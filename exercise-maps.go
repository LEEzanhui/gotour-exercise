package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	str := strings.Fields(s)
	m := make(map[string]int)
	for i := range str {
		m[str[i]]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
