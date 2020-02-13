package main

import "fmt"

func main() {
	// In Go language, a channel is a medium through which a goroutine communicates with another goroutine and this communication is lock-free. Or in other words, a channel is a technique which allows to let one goroutine to send data to another goroutine. By default, a channel is bidirectional which means the go routines can send or receive data through the same channel

	var ch chan int

	fmt.Println(ch)

	ch = make(chan int)
	fmt.Println(ch)

	// shorthand initialization
	c := make(chan int)
	fmt.Println(c)

	// <- channel operator

	// SEND
	c <- 10

	//  RECEIVE
	num := <-c

	fmt.Println(num)

	close(c)

}
