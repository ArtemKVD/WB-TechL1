package main

import "fmt"

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	slice[len(slice)-1] = 0
	return slice[:len(slice)-1]
}

func main() {
	i := 2
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice = remove(slice, i)
	fmt.Println(slice)
}
