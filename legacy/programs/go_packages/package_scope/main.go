package main

import (
	"fmt"

	f "fmt"
)

// import as an alias

// constants, variables or functions are visible to the entire package

func main() {

	fmt.Println("Scope means name visibility")
	sayHello("Wechuli")

	tf := toFahrenheit(boilingPoint)
	fmt.Println("Water Boiling Point in Degrees Fahnrenheint: ", tf)

	f.Println("Hi there")
}
