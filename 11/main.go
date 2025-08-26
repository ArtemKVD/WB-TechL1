package main

import "fmt"

func main() {
	m := make(map[int]int)

	sl1 := []int{1, 2, 3}
	sl2 := []int{2, 3, 4}

	for _, value := range sl1 {
		m[value] += 1
	}

	for _, value := range sl2 {
		m[value] += 1
	}

	sl3 := []int{}

	for i, value := range m {
		if value > 1 {
			sl3 = append(sl3, i)
		}
	}
	fmt.Println(sl3)
}
