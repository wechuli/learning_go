package main

import "fmt"

func main() {
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "John Keller",
		salary: 4000,
		contactInfo: Contact{
			email:   "jkeller@email.com",
			address: "somewhere in Africa",
			phone:   00000323544,
		},
	}

	fmt.Printf("%+v \n", john)

	// accessing field

	fmt.Printf("Employee'e email : %s \n", john.contactInfo.email)

	john.contactInfo.email = "newemail@company.com"

	fmt.Printf("%+v \n", john)
	// Contact

	myContact := Contact{email: "wechuli@email.com", phone: 254658, address: "Nairobi"}

	fmt.Printf("%#+v \n", myContact)
}
