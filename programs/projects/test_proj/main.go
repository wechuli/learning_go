package main

import "fmt"

type Person struct{
	name string
	age int
}

func main() {
	person := Person{"Paul", 23}
	fmt.Println(person)
	fmt.Println("Person")
}
