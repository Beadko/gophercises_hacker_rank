package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	answer := 1
	for _, l := range input {
		str := string(l)
		if strings.ToUpper(str) == str {
			answer++
		}
	}
	fmt.Println(answer)
}
