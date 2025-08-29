package main

import "fmt"

func main() {
	m := make(map[rune]int)
	a := "2146789"
	//a := "221345678"
	for _, val := range a {
		m[val] += 1
	}

	for _, counter := range m {
		if counter > 1 {
			fmt.Print(false)
			return
		}
	}
	fmt.Print(true)
}
