package main

import "fmt"

func main() {

	// must be initialized when declared
	const days int = 7

	fmt.Println("Hi")

	const pi float64 = 3.14
	const secondsInHour = 3600

	duration := 234 // in hours
	_ = duration

	// declaring multiple constants

	const n, m int = 4, 5
	const n1, m1, m2 = 6, 5, 3

	const (
		min1 = 100
		min2 = 200
		min3 = 300
	)
	fmt.Println(min1, min2, min3)

	untypedConstants()

}

func untypedConstants() {
	const a float64 = 5.1 //typed constant
	const b = 6.7         //untyped constants

	const c float64 = a * b // evaluated at compile time
	const str = "Hello" + "Go"
	const d = 5 < 10

	const x = 5
	const y = 2.3 * 5

	fmt.Printf("%T \n", y)
}
