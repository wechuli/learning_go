package main

import "fmt"

func main() {

	type Book struct {
		title  string
		author string
		year   int
	}

	myBook := Book{"The Devine Comedy", "Dante Aligheri", 1320}
	fmt.Printf("%T \n", myBook)
	fmt.Println(myBook)

	myOtherBook := Book{title: "Animal Farm", author: "George Orwell", year: 1945}
	fmt.Println(myOtherBook)
}
