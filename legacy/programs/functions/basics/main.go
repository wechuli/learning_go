package main

import (
	"fmt"
	"math"
)

func f1() {
	fmt.Println("This is f1() function")
}

func f2(a int, b int) {
	fmt.Println("Sum: ", a+b)
}

// shorthand notation of the input parameters
func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

// returning parameters
func f4(a float64) float64 {
	return math.Pow(a, a)
}

// returning multiple params
func f5(a, b int) (int, int) {
	return a + b, a * b
}

// names returned value
func summation(a, b int) (s int, message string) {
	s = a + b
	message = "hello"

	// // naked return
	// return

	return s, message
}

func main() {
	f1()
	f2(34, 65)
	f3(4, 6, 5, 6, 3.2, "dd")
	fmt.Println(f4(5.9))

	sum, product := f5(8, 9)
	fmt.Println(sum, product)

	a, b := summation(78, 34)

	fmt.Println(a, b)
}
