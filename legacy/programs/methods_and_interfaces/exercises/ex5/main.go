package main

import "fmt"

type emptyInterface interface {
}

func main() {
	var empty emptyInterface = 8

	fmt.Println(empty)

	empty = 85.95

	fmt.Println(empty)

	empty = []string{"John", "Peter", "Jess"}

	fmt.Println(empty)

	empty = append(empty.([]string), "Cheese")

	fmt.Println(empty)

}
