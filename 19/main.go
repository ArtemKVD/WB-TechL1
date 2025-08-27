package main

import (
	"fmt"
)

func reverse(str []rune) []rune {
	reversestr := make([]rune, len(str))

	for i := 0; i < len(str); i++ {
		reversestr[len(str)-i-1] = str[i]
	}

	return reversestr
}

func main() {
	var str string
	fmt.Scan(&str)

	a := []rune(str)
	a = reverse(a)

	str = string(a)
	fmt.Println(str)
}
