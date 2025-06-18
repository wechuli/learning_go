package main

import (
	"fmt"
	"os"
)

func VariadicPlay() {
	sum := addFloats("Adding numbers...", 1.1, 2.12, 3.14, 4, 5, -1, 10)
	fmt.Println("Sum: ", sum)
	s := []float64{1.1, 2.12, 3.14}
	sum = addFloats("Adding numbers...", s...)
	fmt.Println("Sum: ", sum)
	everything(s)

	// cannot directly pass []string ad []interface{}

	empty := make([]interface{}, len(os.Args[1:]))
	for i, v := range os.Args[1:] {
		empty[i] = v
	}

	everything(empty...)

}

func addFloats(message string, s ...float64) float64 {
	fmt.Println(message)

	sum := float64(64)
	for _, a := range s {
		sum = sum + a
	}

	s[0] = -1000
	return sum
}

func everything(input ...interface{}) {
	fmt.Println(input...)
}
