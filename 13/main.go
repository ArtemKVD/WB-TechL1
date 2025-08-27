package main

import "fmt"

func main() {
	a := 15
	b := 10

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println(a, " ", b)
}
