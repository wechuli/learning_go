package main

import (
	"fmt"
	"strings"
)

func main() {
	// declare an array
	var numbers [4]int

	fmt.Println(numbers)

	// initializing an array

	var a1 = [4]float64{}
	fmt.Println(a1)

	var a2 = [3]int{-10, 1, 100}
	fmt.Println(a2)

	a3 := [4]string{"Dane", "dsds", "emie52", "ewrew"}
	fmt.Println(a3)

	// ellipsis operator

	a5 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%#v \n", a5)
	fmt.Printf("the length of a5 is %d \n", len(a5))

	// array on multiple lines

	a6 := [...]int{
		1,
		2,
		3,
		4, // ending comma is mandatory for new line declaration
	}

	fmt.Println(a6)

	a6[3] = 900

	fmt.Println(a6)

	// iterating through arrays

	for i, v := range numbers {
		fmt.Println("index: ", i, "value: ", v)
	}

	fmt.Println(strings.Repeat("#", 20))
	for i := 0; i < len(numbers); i++ {
		fmt.Println("index: ", i, "value: ", numbers[i])
	}

	// Mutlidimensional arrays

	balances := [2][3]int{
		{5, 6, 7},
		{8, 9, 10},
	}
	fmt.Println(balances)

	// equality of arrays
	m := [3]int{1, 2, 3}
	n := m
	fmt.Println("n is equal to m: ", n == m)
	n[0] = 78
	fmt.Println("n is equal to m: ", n == m)

	// Arrays with Keyed Elements
	grades := [3]int{
		1: 10,
		0: 5,
		2: 7,
	}
	fmt.Println(grades)

	accounts := [3]int{2: 50}
	fmt.Println(accounts)

	names := [...]string{5: "Dan"}
	fmt.Println(names)

	// if we dont specify key

	cities := [...]string{
		5: "Paris",
		"London",
		1: "NYC",
		"Nairobi",
	}

	fmt.Println(cities)

	weekend := [7]bool{5: true, 6: true}
	fmt.Println(weekend) // => [false false false false false true true]

}
