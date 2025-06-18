package main

import "fmt"

type Book struct {
	title string
	price float64
}

func (b Book) vat() float64 {
	return b.price * 0.09
}

func (b *Book) discount() {
	(*b).price *= 0.9
}

func main() {

	myBook := Book{title: "MiddleMarch", price: 100}
	myBook.vat()
	fmt.Println(myBook)

	myBook.discount()

	fmt.Println(myBook)

}
