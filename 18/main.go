package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
}

func CounterIncrement(c *counter, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	c.count++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	c := &counter{
		count: 0,
	}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go CounterIncrement(c, &mu, &wg)
	}

	wg.Wait()
	fmt.Println(c.count)
}
