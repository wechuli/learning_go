package main

import "fmt"

func main() {

	outer := 19 //no conflict with labels and variables
	_ = outer
	people := [5]string{"Hellen", "Mark", "Brenda", "Antonio", "Michael"}
	friends := [2]string{"Mark", "Mary"}

outer:
	for index, person := range people {
		for _, friend := range friends {
			if person == friend {
				fmt.Printf("Found a friend %q at index %d \n", friend, index)
				break outer
			}
		}
	}
	fmt.Println("Next instruction after the break")

}
