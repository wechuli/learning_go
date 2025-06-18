package main

import "fmt"

func main() {
	a := "This is just a string"
	b := &a
	fmt.Println(a, b)
	fmt.Printf("%T \n", b)

	var x int = 2
	ptr := &x

	fmt.Printf("ptr is of type %T with a value of %v and address %p \n", ptr, ptr, &ptr)

	fmt.Printf("address of x is %p \n", &x)

	// declare a pointer without initializing it

	var ptr1 *float64 // a pointer to that type
	fmt.Println(ptr1)

	p := new(int)

	x = 100
	p = &x

	fmt.Printf("p is of type %T with a value of %v \n", p, p)
	fmt.Printf("address of x is %p \n", &x)

	// deferencning operator
	fmt.Println(*p)
	fmt.Println("*p == x: ", *p == x)

	// pointers to pointers

	val := 5.5
	p1 := &val
	pp1 := &p1

	fmt.Printf("Value of p1: %v, address of p1: %v \n", p1, &p1)
	fmt.Printf("Value of pp1: %v, address of pp1: %v \n", pp1, &pp1)

	fmt.Printf("*pa is %v \n", *p1)
	fmt.Printf("**pp1 is %v \n", *pp1)

	// double deference

	fmt.Printf("**pps is %v \n", **pp1)

	**pp1++
	fmt.Printf("val is %v \n", val)

	// Comparing Pointers

	var p2 *int
	fmt.Printf("%#v \n", p2)

	y := 5
	p2 = &y

	z := 5
	p3 := &z

	p4 := &y

	fmt.Println(p2 == p3)
	fmt.Println(p2 == p4)

}
