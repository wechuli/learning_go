package main

import (
	"fmt"
)

func main() {
	// using the var keyword
	var age int64 = 27
	var firstName string = "Paul"
	var temperature float64

	temperature = 34.38

	// go can type infer
	var girl = "June"

	// if we want to declare a varible without using it

	var schoolName = "fdfd"
	_ = schoolName

	// short hand variable declaring

	goal := "Learn to program in go"

	fmt.Println("Age: ", age)
	fmt.Println("First Name", firstName)
	fmt.Println("Girl Name", girl)
	fmt.Println(temperature)
	fmt.Println(goal)

	// Multiple variable declaration

	car, cost := "Audi", 5000
	fmt.Println(car, cost)

	var course, hours = "Go programming", 14

	fmt.Println(course, hours)

	var opened bool = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	// delcaring mulitple variables
	var (
		salary float64
		name   string
		gender bool
	)
	fmt.Println(salary, name, gender)

	// more on delcaring variables

	var i, j int
	i, j = 5, 8

	fmt.Println(j, i)
	// swap variables

	i, j = j, i

	fmt.Println(j, i)

}
