package main

import (
	"fmt"
	"sync"
)

func FromArrayInChannel(arr []int, channel chan int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for _, value := range arr {
			channel <- value
		}
		close(channel)
	}()
}

func MultipleSend(channelread chan int, channelwrite chan int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for value := range channelread {
			channelwrite <- value * 2
		}
		close(channelwrite)
	}()
}

func StdOutFromChannel(channel chan int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for value := range channel {
			fmt.Println(value)
		}
	}()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	channelread := make(chan int)
	channelwrite := make(chan int)

	FromArrayInChannel(arr, channelread, &wg)
	MultipleSend(channelread, channelwrite, &wg)
	StdOutFromChannel(channelwrite, &wg)

	wg.Wait()
}
