package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

func someFunc() {
	//создали массив на 1024 байта, используем только 100 - утечка памяти
	v := createHugeString(1 << 10)

	justString = v[:100] //присваивает первые 100 символов строки v

	fmt.Println(justString)
}

func main() {
	someFunc()
}
