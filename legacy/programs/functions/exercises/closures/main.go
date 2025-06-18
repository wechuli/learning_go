package main

import "fmt"

func outer(number int) func() int {

	return func() int {
		fmt.Println(number)
		return number
	}
}

// func passFunc(myFunc func()) {

// 	fmt.Println("Hi, I just call functions passed to me")
// 	myFunc()
// }

func anotherFunc(number int) {
	fmt.Println(number)
}
func main() {

	a := anotherFunc

	fmt.Printf("%T \n", a)
	a(856)

	b := outer(85)
	b()

	passFunc(anotherFunc)

}
