package main

import "fmt"

// declare an array

var numbers [10]int

func init() {
	fmt.Println("This is init 1")
}

func init() {
	fmt.Println("This is init 2")

	for i := 0; i < len(numbers); i++ {
		numbers[i] = i * 2
	}
}

func main() {
	fmt.Println("This is main")

	fmt.Printf("%+v \n", numbers)

}
