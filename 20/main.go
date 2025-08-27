package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	s := "fefef a afefgg hthd"
	fmt.Println(reverse(s))
}
