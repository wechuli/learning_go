package main

import "fmt"

func sum(numbers ...int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
}
