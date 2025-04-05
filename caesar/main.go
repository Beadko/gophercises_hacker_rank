package main

import (
	"fmt"
	"strings"
)

func main() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	aL := "abcdefghijklmnopqrstuvwxyz"
	aU := strings.ToUpper(aL)
	r := ""
	for _, c := range input {
		switch {
		case strings.IndexRune(aL, c) >= 0:
			r = r + string(rotate(c, delta, []rune(aL)))
		case strings.IndexRune(aU, c) >= 0:
			r = r + string(rotate(c, delta, []rune(aU)))
		default:
			r = r + string(c)
		}
	}
	fmt.Println(r)
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
