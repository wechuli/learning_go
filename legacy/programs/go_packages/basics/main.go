package main

import (
	"fmt"
	"learning_go/programs/go_packages/mypackages/numbers" // the path should be relative to the GO PATH src directory
)

func main() {

	var x uint = 78
	fmt.Printf("%d is even: %t \n", x, numbers.Even(x))
	fmt.Printf("%d is Prime: %t \n", x, numbers.IsPrime(int(x)))

	fmt.Println("Prime and odd: ", numbers.OddAndPrime(13))

}
