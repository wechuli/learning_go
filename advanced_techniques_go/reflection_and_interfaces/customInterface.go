package main

import (
	"fmt"
	"math"
)

func BasicInterfacePlay() {
	a := circle{R: 1.5}
	fmt.Printf("R %.2f -> Permieter %.3f \n", a.R, a.Perimeter())

	_, ok := interface{}(a).(Shape2D)
	if ok {
		fmt.Println("a is a Shape2D!")
	}

	i := 12
	_, ok = interface{}(i).(Shape2D)
	if ok {
		fmt.Println("i is a Shape2D!")
	} else {
		fmt.Println("i is not a Shape2D!")
	}
}

type Shape2D interface {
	Perimeter() float64
}

type circle struct {
	R float64
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}
