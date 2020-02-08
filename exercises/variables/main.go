package main

import (
	"fmt"
)

func main() {
	var a int = 34
	var b float64 = 45.9
	var c bool = false
	var d string = "this is just a string literal"

	var a1, b1, c1, d1 = 34, 45.9, false, "this is just a string literal"

	a2, b2, c2, d2 := 34, 45.9, false, "this is just a string literal"

	fmt.Println(a, b, c, d)
	fmt.Println(a1, b1, c1, d1)
	fmt.Println(a2, b2, c2, d2)
	// short syntax

	x := 20
	y := 15.5
	z := "Gopher"

	fmt.Println(x, y, z)
}
