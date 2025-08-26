package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(channel <-chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case data := <-channel:
			fmt.Println(data)
		case <-ctx.Done():
			//выходим из горутины если контекст отменен
			return
		}

	}
}

func main() {
	var wg sync.WaitGroup
	var count int
	fmt.Scan(&count)

	channel := make(chan int)

	ctx, cancel := context.WithCancel(context.Background()) // создаём контекст, который может быть отменён
	for i := 0; i < count; i++ {
		wg.Add(1)
		go worker(channel, &wg, ctx)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	data := 0

	go func() {
		<-signalChan
		fmt.Println("Shutting down")
		cancel()       //отменяем контекст
		close(channel) // закрываем канал
	}()
	go func() {
		for {
			channel <- data
			data += 10
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}
