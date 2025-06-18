package main

import (
	"fmt"
)

func Userinput(){
	fmt.Println("Please give me your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is ", name)
}