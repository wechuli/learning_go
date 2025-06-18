package main

import "fmt"

func Numerics() {
	// complex

	c1 := 12 + 1i
	c2 := complex(5, 7)
	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of c2: %T\n", c2)

	// integer variables

	

	x := 12
	k := 5

	fmt.Println(x, k)
	fmt.Printf("Type of x: %T\n", x)

	div := x / k
	fmt.Println("div", div)
	fmt.Printf("Type od div: %T\n", div)

	// floats

	var m, n float64

	m = 1.123
	fmt.Println("m,n: ", m, n)

	y := 4 / 2.3
	fmt.Println("y:", y)

	divFloat := float64(x) / float64(y)
	fmt.Println("divFloat", divFloat)
	fmt.Printf("Type of divFloat: %T\n", divFloat)

}
