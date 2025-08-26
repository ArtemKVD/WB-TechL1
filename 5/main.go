package main

import (
	"fmt"
	"time"
)

func main() {
	timerchan := time.After(10 * time.Second)
	counter := 0
	channel := make(chan int)
	go func() {
		for {
			select {
			case <-timerchan:
				fmt.Println("timer end")
				return
			default:
				channel <- counter
				time.Sleep(time.Second)
				counter++
			}
		}
	}()

	for {
		select {
		case <-timerchan:
			fmt.Println("timer end")
			return
		case data := <-channel:
			fmt.Println(data)
		}
	}
}
