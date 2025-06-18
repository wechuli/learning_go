package main

import (
	"fmt"
	"os"
	"strconv"
)

type ar2x2 [2][2]int

// traditional Add() function

func TypeMethodsPlay() {
	if len(os.Args) != 9 {
		fmt.Println("Need 8 integeres")
		return
	}

	k := [8]int{}

	for index, i := range os.Args[1:] {
		v, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println(err)
			return
		}
		k[index] = v
	}

	a := ar2x2{{k[0], k[1]}, {k[2], k[3]}}
	b := ar2x2{{k[4], k[5]}, {k[6], k[7]}}

	fmt.Println("Traditional a+b:", Add(a, b))
	a.Add(b)
	fmt.Println("a+b", a)
	a.Subtract(a)
	fmt.Println("a-a:", a)

	a = ar2x2{{k[0], k[1]}, {k[2], k[3]}}

	a.Multiply(b)
	fmt.Println("a*b: ", a)

	a = ar2x2{{k[0], k[1]}, {k[2], k[3]}}
	b.Multiply(a)
	fmt.Println("b*a:", b)
}

func Add(a, b ar2x2) ar2x2 {
	c := ar2x2{}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c

}

func (a *ar2x2) Add(b ar2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] = a[i][j] + b[i][j]
		}
	}
}

func (a *ar2x2) Subtract(b ar2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] = a[i][j] - b[i][j]
		}
	}
}

func (a *ar2x2) Multiply(b ar2x2) {
	result := ar2x2{} // Create a temporary matrix to store the result
	result[0][0] = a[0][0]*b[0][0] + a[0][1]*b[1][0]
	result[1][0] = a[1][0]*b[0][0] + a[1][1]*b[1][0]
	result[0][1] = a[0][0]*b[0][1] + a[0][1]*b[1][1]
	result[1][1] = a[1][0]*b[0][1] + a[1][1]*b[1][1]

	*a = result // Update the original matrix with the result
}
