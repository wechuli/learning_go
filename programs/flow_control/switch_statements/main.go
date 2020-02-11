package main

import (
	"fmt"
	"time"
)

func main() {
	language := "golang"

	switch language {
	case "Python":
		fmt.Println("You are learning Python! You don't use curly braces bu indentation!")
		// break - you don't need an explicit break
	case "Go", "golang":
		fmt.Println("Good, go for go You atr using curly braces")
	default:
		fmt.Println("Any other programming language is good")

	}
	n := 5
	switch true { // we can use boolean as the expression
	case n%2 == 0:
		fmt.Println("Even number")
	case n%2 != 0:
		fmt.Println("Odd number")

	}

	hour := time.Now().Hour()
	fmt.Println(hour)

	switch { // this is equivalent to true
	case hour < 12:
		fmt.Println("Good Morning !")
	case hour < 17:
		fmt.Println("Good Afternoon !")
	default:
		fmt.Println("Good Evening !")

	}
}
