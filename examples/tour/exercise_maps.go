// Exercise: Loops and Functions

package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(sentence string) map[string]int {
	m := make(map[string]int)
	s := strings.Fields(sentence)
	for _, word := range s {
		m[word] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
