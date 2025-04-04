package main

import "fmt"

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
	idx := -1
	for i, r := range k {
		if r == s {
			idx = i
			break
		}
	}
	if idx < 0 {
		fmt.Printf("Error: could not find the expected symbol %v\n", s)
		return s
	}
	for i := 0; i < delta; i++ {
		idx++
		if idx >= len(k) {
			idx = 0
		}
	}
	return k[idx]
}
