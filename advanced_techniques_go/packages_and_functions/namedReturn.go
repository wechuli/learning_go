package main

import (
	"fmt"
	"os"
	"strconv"
)

func NamedReturnPlay() {
	arguments := os.Args
	if len(arguments) < 3 {
		fmt.Println("The program needs at least 2 arguments!")
		return
	}

	a1, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("Must use integeres")
		return
	}
	a2, err := strconv.Atoi(arguments[2])
	if err != nil {
		fmt.Println("Must use integeres")
		return
	}

	fmt.Println(minMax(a1, a2))
	mi, ma := minMax(a1, a2)
	fmt.Println(mi, ma)
}

func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
		return min, max
	}
	min = x
	max = y
	return
}
