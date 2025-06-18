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

	lastBook := Book{title: "Anna Karenina"}

	// retriving value

	fmt.Println(lastBook.title)

	lastBook.author = "Leo Tolstoy"
	lastBook.year = 1878

	fmt.Printf("%+v \n", lastBook)

	// comparing struct

	aBook := Book{title: "Animal Farm", author: "George Orwell", year: 1945}

	fmt.Println(aBook == myOtherBook)

	// making a copy of a struct

	copyBook := aBook

	copyBook.year = 2020

	fmt.Println(copyBook)
	fmt.Println(aBook)
}
