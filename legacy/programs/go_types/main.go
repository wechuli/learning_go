package main

import (
	"fmt"
)

func main() {
	// INT TYPE
	var i1 int8 = 100

	fmt.Printf("%T \n", i1)

	var i2 uint16 = 65535
	fmt.Printf("%T \n", i2)

	// floats

	var f1, f2, f3 float64 = 1.1, -.2, 5

	fmt.Printf("%T %T %T \n", f1, f2, f3)

	// Rune type
	var r rune = 'f'
	fmt.Printf("%T \n", r)
	fmt.Println(r) //print the int value
	fmt.Printf("%x \n", r)

	// Bool type

	var b bool = true
	fmt.Println(b)

	// String type

	var myPoem string = "This is a poem of mine"
	fmt.Println(myPoem)

	compositeType()

}

func compositeType() {

	//Array Tyoe

	var numbers = [4]int{4, 5, -9, 100}
	fmt.Println(numbers)
	fmt.Printf("%T", numbers)

	// Slice Type

	var cities = []string{"Nairobi", "Kakamega", "Dar"}
	fmt.Printf("%T \n", cities)
	fmt.Println(cities[2])

	// Map
	balances := map[string]float64{
		"USD": 2343.434,
		"EUR": 323,
	}
	fmt.Println(balances)
	fmt.Println(balances["USD"])

	// Struct

	type Person struct {
		name string
		age  int
	}

	var you Person

	fmt.Printf("%T \n", you)
	fmt.Println(you)

	// Pointer Type
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with the value of %v \n", ptr, ptr)

	// Fucntion type

}
