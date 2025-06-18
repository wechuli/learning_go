package main

import "fmt"

func main() {
	type age int        // int is its underlying type
	type oldAge age     // int is its underlying type
	type veryOldAge age // int is its underlying type

	type speed uint

	var s1 speed = 10
	var s2 speed = 20

	fmt.Println(s1 + s2)

	var x uint

	// you cannot do this
	// x = s1

	s1 = speed(x)
	fmt.Printf("%T \n", s1)

	type km float64

	type mile float64

	var parisToLondom km = 465
	var distanceInMiles mile

	// can be converting if the underlying type is the same
	distanceInMiles = mile(parisToLondom) / 0.621
	fmt.Println(distanceInMiles)
}

func classNotes() {
	// aliases
	// declaring a variable of type uint8
	var a uint8 = 10
	var b byte // byte is an alias to uit8

	// even though they have different names, byte and uit8 are the same type because they are aliases
	b = a // no error
	_ = b

	// declaring a new alias named second for uint
	// type alias_name = type_name
	type second = uint

	var hour second = 3600
	fmt.Printf("hour type: %T\n", hour) // => hour type: uint

	//no need to convert operations (same type)
	fmt.Printf("Minutes in an hour: %d\n", hour/60) // => Minutes in an hour: 60[



	
	]
}
