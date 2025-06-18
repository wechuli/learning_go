package main

import (
	"fmt"
	"os"
)

func main() {

	Which()

	// control_flow
	//ControlFlow(os.Args)

	// loops
	// Loops()
	//   Userinput()

	//Errvars()

	//for i:=0;i<4;i++{
	//	go MyPrint(i,5)
	//}
	//time.Sleep(time.Second)
	// Which()
	// Logger()

	//Logger2()

	//CustomLogger()

	//fmt.Println("Returned")

}

func runPhoneBook() {
	InitializeData()
	arguments := os.Args

	if len(arguments) == 1 {
		exe := arguments[0]
		fmt.Printf("Usage: %s search|list <arguments> \n", exe)
		return
	}

	term := arguments[1]

	switch term {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := Search(arguments[2])

		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		fmt.Println(*result)
	case "list":
		List()

	default:
		fmt.Println("Not a valid option")

	}
}
