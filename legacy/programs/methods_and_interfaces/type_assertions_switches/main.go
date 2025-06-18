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

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func printShape(s shape) {
	fmt.Printf("Shape: %#v \n", s)
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perimeter())
}

func main() {

	var s shape = circle{radius: 2.5} // we can only access the methods defined

	fmt.Printf("%T \n", s)

	// type assertion
	ball, ok := s.(circle)

	if ok == true {
		fmt.Printf("Ball volume: %v \n", ball.volume())
	}

	// type switches

	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type \n", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type \n", value)

	}


}
