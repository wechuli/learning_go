package main

import (
	"fmt"
	"strings"
)

func factorial(n int, c chan int) {

	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}

	c <- f
}

func main() {

	ch := make(chan int)

	defer close(ch)

	go factorial(5, ch)

	f := <-ch

	fmt.Println(f)

	for i := 1; i < 20; i++ {
		go factorial(i, ch)
		f := <-ch
		fmt.Printf("Factorial of %v : %v \n", i, f)
	}

	fmt.Println(strings.Repeat("#", 10))

	for i := 0; i <= 15; i++ {
		go func(n int, c chan int) {
			f := 1
			for i := 2; i <= n; i++ {
				f *= i
			}

			c <- f
		}(i, ch)

		f := <-ch
		fmt.Printf("Factorial of %v : %v \n", i, f)
	}
}
