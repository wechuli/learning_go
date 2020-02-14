package main

import (
	"fmt"
	"math"
)

func power(number int, ch chan int) {
	ch <- int(math.Pow(float64(number), 2))
}

func main() {
	ch := make(chan int, 50)
	defer close(ch)

	for i := 1; i <= 50; i++ {
		go power(i, ch)

		fmt.Printf("The square of %d is %d \n", i, <-ch)

	}
}
