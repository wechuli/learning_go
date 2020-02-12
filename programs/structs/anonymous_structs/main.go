package main

import "fmt"

func main() {

	// anonymous structs
	diana := struct {
		firstName, lastName string
		age                 int
	}{firstName: "Diana",
		lastName: "Muller",
		age:      30}

	fmt.Printf("%#v \n", diana)
	fmt.Printf("Diana's Age: %d \n", diana.age)

	// anonymous fields

	type Book struct {
		string
		float64
		bool
	}

	book1 := Book{"1984 by George Orwell", 10.2, false}

	fmt.Printf("%#v \n", book1)

	// mixing named fields and anonymous fields

	type Employee struct {
		name   string
		salary int
		bool
	}

	john := Employee{"John", 40000, false}
	fmt.Printf("%#v \n", john)
	john.bool = true
	fmt.Printf("%#v \n", john)

}
