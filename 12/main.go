package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]bool)
	for _, value := range arr {
		m[value] = true
	}
	arr2 := []string{}
	for key, _ := range m {
		arr2 = append(arr2, key)
	}

	fmt.Println(arr2)
}
