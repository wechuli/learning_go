package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("45")

	if err != nil {
		println(err)
	} else {
		fmt.Println(i)
	}

	if i, err := strconv.Atoi("20"); err == nil {
		fmt.Println("No error, i is: ", i)
	} else {
		fmt.Println(err)
		// handle the error
	}

	if args := os.Args; len(args) != 2 {
		fmt.Println("One argument is required")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("The argument must be an integer! Error: ", err)
	} else {

		fmt.Printf("%T \n", km)
		fmt.Printf("%d km in miles in %v \n", km, float64(km)*0.621)
	}

}
