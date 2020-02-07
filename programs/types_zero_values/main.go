package main

import (
	"fmt"
)

func main() {

	// this is a comment
	var (
		value int
		price float64
		name  string
		done  bool
	)

	/*
		This is a multiline comment
		Deal with it

	*/

	fmt.Println(value, price, name, done)

	var firstName string = "Paul" //this is an inline comment

	fmt.Println(firstName)

	maxValue := 1000  // recommended
	max_values := 100 // Not recommnded

	writeToDB := true  //ok, idiomatic
	writeToDb := false // not okay

	fmt.Println(maxValue, max_values, writeToDB, writeToDb)

	// use camelCase instead of snake_case

}

func writeToFile() { //recommended

}
