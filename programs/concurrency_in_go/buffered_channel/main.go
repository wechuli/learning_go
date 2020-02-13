package main

import (
	"fmt"
	"time"
)

func main() {
	buffCh := make(chan int, 3)

	go func(c chan int) {

		for i := 0; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel \n", i)
			buffCh <- i
			fmt.Printf("func goroutine #%d after sending data into the channel \n", i)
		}
		close(c)

	}(buffCh)

	fmt.Println("main goroutine sleeps for 2 seconds")

	time.Sleep(time.Second * 2)

	for v := range buffCh {
		fmt.Println("Main goroutine received value from channel.", v)
	}

	fmt.Println(<-buffCh)

	// c <- buffCh // panic

}
