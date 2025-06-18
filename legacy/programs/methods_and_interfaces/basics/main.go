package main

import (
	"fmt"
	"time"
)

type names []string

func (n names) print() {
	for i, name := range n {
		fmt.Println(i, name)
	}
}

func main() {
	const day = 24 * time.Hour
	fmt.Printf("%T \n", day)

	seconds := day.Seconds()
	fmt.Printf("%T \n", seconds)
	fmt.Printf("Seconds in a day: %v \n", seconds)

	friends := names{"Dan", "Mary", "John"}

	friends.print()

	// another way to call the method - method belongs to a type, function belongs to a package

	names.print(friends)

}
