package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func printShape(s shape) {
	fmt.Printf("Shape: %#v \n", s)
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perimeter())
}

func main() {
	var s shape
	fmt.Printf("%T \n", s)

	ball := circle{radius: 2.5}
	s = ball // the interface can hold the value of a conrete type

	fmt.Printf("%T \n", s)

	printShape(s)

	room := rectangle{width: 2, height: 3}
	s = room
	fmt.Printf("Type of s: %T \n", s)

}
