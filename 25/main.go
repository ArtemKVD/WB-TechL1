package main

import (
	"fmt"
	"time"
)

func mySleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("start")
	mySleep(2 * time.Second)
	fmt.Println("finish")
}
