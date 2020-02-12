package main

import "fmt"

type Person struct {
	name      string
	age       int
	favColors []string
}

func main() {
	me := Person{name: "Paul", age: 27, favColors: []string{"Green", "Teal", "Maroon"}}
	you := Person{name: "June", age: 24, favColors: []string{"Purple", "baige"}}

	fmt.Printf("%T \n", me)
	fmt.Printf("%+v \n", me)
	fmt.Printf("%+v \n", you)

	me.name = "Wechuli"

	fmt.Printf("%+v \n", me)

	var youFavColors = you.favColors

	fmt.Printf("%T \n", youFavColors)
	fmt.Printf("%+v \n", youFavColors)

	you.favColors = append(you.favColors, "black", "brown")

	fmt.Printf("%+v \n", you)

	for index, color := range me.favColors {
		fmt.Println(index, color)
	}
}
