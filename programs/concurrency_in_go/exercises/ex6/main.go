package main

import (
	"fmt"
	"strings"
)

func main() {
	ch := make(chan string)

	defer close(ch)

	go func(dummy string, ch chan string) {
		ch <- strings.Repeat(dummy, 2)
	}("Hello", ch)

	fmt.Println("This was sent back from the channel: ", <-ch)
}
