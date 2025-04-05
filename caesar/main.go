package main

import (
	"fmt"
	"strings"
)

func main() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%d\n", &delta)
	fmt.Scanf("%d\n", &input)

	a := []rune("abcdefghijklmnopqrstuvwxyz")
	newRune := rotate('m', 2, a)
	fmt.Println(string(newRune))
}

func rotate(s rune, delta int, k []rune) rune {
	idx := strings.IndexRune(string(k), s)
	if idx < 0 {
		fmt.Printf("Error: could not find the expected symbol %v\n", s)
		return s
	}
	idx = (idx + delta) % len(k)
	return k[idx]
}
