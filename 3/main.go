package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(channel <-chan int, wg *sync.WaitGroup) {
	for data := range channel {
		fmt.Println(data)
	}
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var count int
	fmt.Scan(&count)

	channel := make(chan int)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go worker(channel, &wg)
	}

	data := 0
	for {
		channel <- data
		data += 10
		time.Sleep(time.Second)
	}
}
