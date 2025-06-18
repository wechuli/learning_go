package main

import "fmt"

type Credential struct {
	SecretValue string
}

type Entry struct {
	F1 int
	F2 string
	F3 Credential
}

func TypeSwitchPlay() {
	A := Entry{100, "F2", Credential{"mypassword"}}
	Teststruct(A)
	Teststruct(A.F3)
	Teststruct("A string")

	Learn(12.23)
	Learn('€')
}

func Teststruct(x interface{}) {
	// type switch

	switch T := x.(type) {
	case Credential:
		fmt.Println("Secret type")
	case Entry:
		fmt.Println("Entry type")
	default:
		fmt.Printf("Not supported type: %T\n", T)
	}
}

func Learn(x interface{}) {
	switch T := x.(type) {
	default:
		fmt.Printf("Data type: %T\n", T)
	}
}
