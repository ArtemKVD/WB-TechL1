package main

import (
	"fmt"
	"sync"
)

func Write(m map[int]int, mu *sync.Mutex, key int, value int) {
	mu.Lock()
	defer mu.Unlock()
	m[key] = value
}

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	m := make(map[int]int)

	wg.Add(30)

	for i := 0; i < 30; i++ {
		go func(i int) {
			Write(m, &mu, i, i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	for key, value := range m {
		fmt.Println(key, " ", value)
	}

}
