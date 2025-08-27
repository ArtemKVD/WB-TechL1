package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("10000000000000000000", 10)
	b.SetString("100000000000000000000000", 10)

	result := new(big.Int)

	result.Add(a, b)
	fmt.Println(result.String())

	result.Sub(a, b)
	fmt.Println(result.String())

	result.Mul(a, b)
	fmt.Println(result.String())

	result.Div(b, a)
	fmt.Println(result.String())
}
