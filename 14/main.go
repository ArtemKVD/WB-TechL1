package main

import "fmt"

func CheckType(a interface{}) {
	switch t := a.(type) {
	case int:
		fmt.Println("int", t)
	case string:
		fmt.Println("string", t)
	case bool:
		fmt.Println("bool", t)
	case chan int:
		fmt.Println("chan int")
	case chan string:
		fmt.Println("chan string")
	case chan bool:
		fmt.Println("chan bool")
	}
}
func main() {
	CheckType(123)
	CheckType("1234")
	CheckType(true)
	chanint := make(chan int)
	CheckType(chanint)
	chanbool := make(chan bool)
	CheckType(chanbool)
	chanstring := make(chan string)
	CheckType(chanstring)
}
