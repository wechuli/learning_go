package main

import (
	"fmt"
)

func main() {
	name := "Paul"
	fmt.Println("Hello,playground", name, "hi there")

	a, b := 4, 6

	fmt.Println("Sum: ", a+b, "Mean Value:", (a+b)/2)

	// using printf

	fmt.Printf("Your age is %d \n", a)
	fmt.Printf("x is %d, y is %f \n", 5, 7.8)
	fmt.Printf("He says: \"Hello Go!\"  \n")

	var figure string = "Circle"
	var radius int64 = 5
	pi := 3.14149

	fmt.Printf("Radius is %d \n", radius)
	fmt.Printf("Radius is %+d \n", radius)

	fmt.Printf("Pi constant is %f\n", pi)
	fmt.Printf("The diameter of a %s with a Radius of %d is %f \n", figure, radius, float64(radius)*2*pi)

	// %q for quoted string

	fmt.Printf("This is a %q", figure)

	// %v -> replaced by any type

	fmt.Printf("The diameter of a %v with a Radius of %v is %v \n", figure, radius, float64(radius)*2*pi)

	// %T - > print out the type

	fmt.Printf("figure is of type %T \n", figure)

	// %t -> format a boolean value as true or false
	closed := true
	fmt.Printf("File closed: %t \n", closed)

	// %b - base2

	fmt.Printf("%b \n", 504)

	// %x - base 16

	fmt.Printf("%x", 15)

	const pi2 float64 = 3.14159265359
	fmt.Printf("pi is %.4f \n", pi2)

}
