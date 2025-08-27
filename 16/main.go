package main

import "fmt"

func quickSort(arr []int) []int {
	var sorted []int
	if len(arr) < 2 {
		return arr
	}

	middle := arr[len(arr)/2]

	var l, m, r []int

	for _, item := range arr {
		switch {
		case item < middle:
			l = append(l, item)
		case item == middle:
			m = append(m, item)
		case item > middle:
			r = append(r, item)
		}
	}

	l = quickSort(l)
	r = quickSort(r)

	sorted = append(sorted, l...)
	sorted = append(sorted, m...)
	sorted = append(sorted, r...)

	return sorted
}

func main() {
	arr := []int{10, 20, 30, 40, 50, 1, 4, 5, 6, 8, 111}
	fmt.Println(quickSort(arr))
}
