package main

import "fmt"

func main() {
	var ch chan float64

	// declaring and initializing a Receive only channel
	ch2 := make(<-chan rune)

	// declaring and initializing a send-only channel
	ch3 := make(chan<- rune)

	ch4 := make(chan rune, 10)

	fmt.Printf("%T, %T, %T, %T \n", ch, ch2, ch3, ch4)
}
