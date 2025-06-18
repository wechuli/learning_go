package main

import (
	"fmt"
	"os"
	"strconv"
)

func ControlFlow(args []string){
	if (len(args) < 2){
		fmt.Println("Please provide a command line argument")
		return
	}
	
	argument := os.Args[1]

	value,err := strconv.Atoi(argument)

	if err != nil {
		fmt.Println("Cannot convert to int: ", argument)
		return
	}

	switch(value){
    case 0:
		fmt.Println("Zero!")
	case 1:
		fmt.Println("One!")
	case 2,3,4:
		fmt.Println("2 or 3 or 4")
	default:
		fmt.Println("Value: ", argument)

	}

	switch {
	case value == 0:
		fmt.Println("Zero!")
	case value > 0:
		fmt.Println("Positive integer")
	case value < 0:
		fmt.Println("Negatove integer")
	default:
		fmt.Println("This should not happen", value)
	}
}
