package main

import "fmt"

func main() {
	num := int64(5)
	var i int
	fmt.Scan(&i)
	num = num ^ (1 << i)
	fmt.Println(num)
	num = num ^ (1 << i)
	fmt.Println(num)
}
