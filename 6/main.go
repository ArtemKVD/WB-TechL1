package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func Exit() {
	go func() {
		counter := 0
		for counter < 5 {
			fmt.Println(counter)
			counter++
			time.Sleep(time.Second)
		}
		fmt.Println("горутина завершилась")
	}()
	time.Sleep(5 * time.Second)
}

func ExitByChannel() {
	Channel := make(chan bool)

	go func() {
		for {
			select {
			case <-Channel:
				fmt.Println("сигнал остановки получен")
				return
			default:
				fmt.Println("горутина работает")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	Channel <- true
	time.Sleep(time.Second)
}

func ExitByCloseChannel() {
	channel := make(chan int)

	go func() {
		for {
			data, ok := <-channel
			if !ok {
				fmt.Println("канал закрыт, завершаем работу")
				return
			} else {
				fmt.Println(data)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		channel <- i
		time.Sleep(time.Second)
	}

	close(channel)
	time.Sleep(time.Second)
}

func ExitByContextCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	counter := 0
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("контекст отменён, завершаем работу")
				return
			default:
				counter++
				fmt.Println(counter)
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(time.Second)
}

func ExitByContextTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	counter := 0
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("таймаут контекста истёк, завершаем работу")
				return
			default:
				counter++
				fmt.Println(counter)
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
}

func ExitByGoExit() {
	go func() {
		fmt.Println("Работаем")
		time.Sleep(time.Second)
		fmt.Println("Вызываем Goexit")
		runtime.Goexit()
	}()
	time.Sleep(5 * time.Second)
}

func main() {
	Exit() // выход по условию
	fmt.Println("--------------------------")
	ExitByChannel() // выход через уведомления
	fmt.Println("--------------------------")
	ExitByCloseChannel() // выход через закрытие канала
	fmt.Println("--------------------------")
	ExitByContextCancel() // выход через отмену контекста
	fmt.Println("--------------------------")
	ExitByContextTimeout() // выход по истечении таймаута контекста
	fmt.Println("--------------------------")
	ExitByGoExit() // выход по runtime.Goexit()

}
