package main

import "fmt"

func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.4
	name = "Mobile Phone"
	sold = false

}

func changeValuesPointer(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 3
	*price = 500.4
	*name = "Mobile Phone"
	*sold = false

}

type Product struct {
	price       float64
	productName string
}

func changeProduct(product Product) {
	product.price = 100
	product.productName = "Bicycle"
}

func changeProductByPointer(product *Product) {
	(*product).price = 1000

	product.productName = "Bicycle"
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 30
	m["b"] = 40
	m["c"] = 50
}

func main() {
	quantity, price, name, sold := 5, 300.4, "Laptop", true
	fmt.Println("Before calling changeValues(): ", quantity, price, name, sold)
	changeValues(quantity, price, name, sold)

	fmt.Println("After calling changeValues(): ", quantity, price, name, sold)

	fmt.Println("Before calling changeValuesPointer(): ", quantity, price, name, sold)
	changeValuesPointer(&quantity, &price, &name, &sold)

	fmt.Println("After calling changeValuesPointer(): ", quantity, price, name, sold)

	// structs
	gift := Product{
		price:       100,
		productName: "Car",
	}
	changeProduct(gift)
	fmt.Println(gift)

	changeProductByPointer(&gift)
	fmt.Println(gift)

	// slicing

	prices := []int{1, 2, 3}
	changeSlice(prices)
	fmt.Println(prices)

	// maps

	myMap := map[string]int{"a": 1, "b": 2}
	changeMap(myMap)
	fmt.Println(myMap)

}
