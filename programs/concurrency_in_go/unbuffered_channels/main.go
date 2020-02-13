package main

import (
	"fmt"
	"time"
)

func main() {
	// unbuffered channel

	c1 := make(chan int)

	// buffered channel
	c2 := make(chan int, 3)
	_ = c2

	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")

	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data")

	d := <-c1
	fmt.Println("Main goroutine received data.", d)

	time.Sleep(time.Second)
}
