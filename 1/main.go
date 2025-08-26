package main

import "fmt"

type Human struct {
	Age    int
	Height int
	Weight int
}

type Action struct {
	Human
	Info string
}

func (h *Human) speak() {
	fmt.Println("Hello!")
}

func (a *Action) sayinfo() {
	fmt.Println(a.Info)
}

func main() {
	human := &Human{
		Age:    20,
		Height: 200,
		Weight: 90,
	}

	action := &Action{
		Human: *human,
		Info:  "Working",
	}

	human.speak()
	action.speak()
	action.sayinfo()
}
