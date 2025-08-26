package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	nums := []int{2, 4, 6, 8, 10}

	for i, num := range nums {
		wg.Add(1)
		go func() {
			nums[i] = num * num
			wg.Done()
		}()
	}
	wg.Wait()

	for _, num := range nums {
		fmt.Println(num)
	}
}
