package main

import (
	"fmt"
	f "fmt"
)

// permitted

const done = false // package scoped constant

func main() {
	var task = "Running" // local (block scoped)
	fmt.Println(task, done)
	f.Println("Bye bye")

	const done = true // permitted since the earlier defined const is package scoped

	fmt.Println(done)
	f1()
}

func f1() {
	fmt.Printf("Done in f1(): %v \n", done) //this is done from name
}
