package main

import (
	"fmt"
)

func main() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	var r []rune
	for _, c := range input {
		r = append(r, cipher(c, delta))
	}
	fmt.Println(string(r))
}

func cipher(r rune, delta int) rune {
	if r >= 'A' && r <= 'Z' {
		return rotate(r, 'A', delta)
	}
	if r >= 'a' && r <= 'z' {
		return rotate(r, 'a', delta)
	}
	return r
}

func rotate(r rune, base, delta int) rune {
	tmp := int(r) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}
