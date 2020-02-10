package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Question 4

	// Create a Go program that reads some numbers from the command line and then calculates the sum and the product of all the numbers given at command line.

	// The user should give between 2 and 10 numbers.

	enteredNumbers := []float64{}
	var errorWithArgs = false

	if len(os.Args) <= 2 || len(os.Args) > 10 {
		log.Fatal("Please run the program with arguments (2-10 numbers)!")
	} else {
		for i := 1; i < len(os.Args); i++ {
			result, err := strconv.ParseFloat(os.Args[i], 64)
			if err == nil {
				enteredNumbers = append(enteredNumbers, result)

			} else {
				errorWithArgs = true
				break
			}
		}
	}

	if errorWithArgs {
		fmt.Println("There was an error with the arguments")
	} else {
		fmt.Printf("The product of the entered numbers is: %.3f \n", getProduct(enteredNumbers))
		fmt.Printf("The sum of the entered numbers is: %.3f \n", getSum(enteredNumbers))
	}
}

func getProduct(numbers []float64) float64 {

	var product float64 = 1
	for _, number := range numbers {
		product *= number
	}
	return product
}

func getSum(numbers []float64) float64 {
	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}
