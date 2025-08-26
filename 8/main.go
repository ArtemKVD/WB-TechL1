package main

import "fmt"

func main() {
	num := int64(5) // 101
	var i int
	fmt.Scan(&i)
	num = num ^ (1 << i)
	fmt.Println(num) // 100
	num = num ^ (1 << i)
	fmt.Println(num) // 101
}
