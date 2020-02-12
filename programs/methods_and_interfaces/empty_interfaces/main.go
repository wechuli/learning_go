package main

import "fmt"

// any go type satisfies this interface
type emptyInterface interface {
}

type person struct {
	info interface{}
}

func main() {
	var empty interface{}

	empty = 5
	fmt.Printf("%T \n", empty)
	fmt.Println(empty)

	empty = "Go"
	fmt.Printf("%T \n", empty)
	fmt.Println(empty)

	empty = []int{4, 5, 6}
	fmt.Println(empty)

	fmt.Println(len(empty.([]int)))

	you := person{}
	you.info = "Your name"
	fmt.Println(you.info)

	you.info = 40

	fmt.Println(you.info)

}
