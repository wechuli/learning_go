package main

import (
	"fmt"
	"math"
)

func cube(side float64) float64 {
	return math.Pow(side, 3)
}

func recursionSum(number uint) (recurs uint, sum uint) {

	recurs = 1
	sum = 0
	for i := number; i > 0; i-- {
		recurs *= i
		sum += i
	}

	return recurs, sum
}

func main() {
	fmt.Println(cube(3))

	recurs, sum := recursionSum(4)

	println(recurs, sum)

	recurs, sum = recursionSum(10)

	println(recurs, sum)
}
