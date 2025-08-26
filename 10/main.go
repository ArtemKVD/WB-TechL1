package main

import "fmt"

func main() {
	m := make(map[int][]float32)
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for _, value := range arr {
		key := int(value) - int(value)%10
		m[key] = append(m[key], value)
	}

	for key, value := range m {
		fmt.Println(key, " ", value)
	}
}
