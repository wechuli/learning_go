package main

import "fmt"

func swapValues(first *float64, second *float64) {
	*first, *second = *second, *first
}

func main() {
	x := 10.10

	ptr := &x

	fmt.Printf("Address of x: %p \n", &x)

	fmt.Printf("Type of ptr: %T \n", ptr)
	fmt.Printf("Value of ptr: %+#v \n", ptr)

	fmt.Printf("Value of x through the pointer: %v \n", *ptr)

	// Question 2

	x, y := 10., 2.

	// fmt.Printf("%T,%T \n", x, y)
	ptrx, ptry := &x, &y

	z := *ptrx / *ptry

	fmt.Println(z)

	// function that swaps values

	value1, value2 := 5.5, 8.8

	fmt.Println("Before Running Swaping Function")
	fmt.Printf("Value 1: %v, Value 2: %v \n", value1, value2)

	swapValues(&value1, &value2)
	fmt.Println("After Running Swaping Function")
	fmt.Printf("Value 1: %v, Value 2: %v \n", value1, value2)

}
