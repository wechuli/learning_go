package main

import "fmt"

func outer(number int) func() int {

	return func() int {
		fmt.Println(number)
		return number
	}
}

func anotherFunc(number int) {
	fmt.Println(number)
}
func main() {

	a := anotherFunc

	fmt.Printf("%T \n", a)
	a(856)

	b := outer(85)
	b()

}
