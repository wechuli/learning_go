package main

import "fmt"

func main() {

	var employees map[string]string

	fmt.Printf("%#v \n", employees)

	fmt.Printf("No of pairs: %d \n", len(employees))

	fmt.Printf("The value for the %q is %q \n", "Dan", employees["Dan"]) // if an element doesn't exist, the zero value is returned

	var accounts map[string]float64

	fmt.Printf("%#v \n", accounts["0x323"])

	// only comparable types can be used as keys

	var m1 map[[5]int]string
	_ = m1

	//employees["Dan"] = "Programmer" //we cant do this

	people := map[string]float64{} //initialized map

	people["John"] = 21.4
	people["Mary"] = 10
	fmt.Println(people)

	// create an initialized but empty map

	map1 := make(map[int]int)
	map1[4] = 8

	// using map literal

	balances := map[string]float64{
		"USD": 323.11,
		"EUR": 432.14,
		"KES": 85, //mandatory comma for multiple line map declaration

	}
	fmt.Println(balances)

	// updating the value

	balances["USD"] = 100

	fmt.Println(balances)

	// disinguishing from sero value to

	v, ok := balances["RON"]

	if ok {
		fmt.Println("The RON balance is: ", v)
	} else {
		fmt.Println("The RON key doesn't exist in the map!")
	}

	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v \n", k, v)
	}

	// deleting values

	delete(balances, "USD")
	fmt.Println(balances)

	// Comparing maps

	a := map[string]string{"A": "X"}
	b := map[string]string{"B": "Y"}

	// fmt.Println(a == b) //invalid
	aString := fmt.Sprintf("%s", a)
	bString := fmt.Sprintf("%s", b)

	fmt.Println(aString == bString)

	a["B"] = "Y"
	b["A"] = "X"

	fmt.Println(a)
	fmt.Println(b)
	aString2 := fmt.Sprintf("%s", a)
	bString2 := fmt.Sprintf("%s", b)

	fmt.Println(aString2 == bString2)

	friends := map[string]int{"Dan": 40, "Maria": 25}

	neighbors := friends

	friends["Dan"] = 50

	fmt.Println(neighbors)

	// copy map

	persons := make(map[string]int)

	for k, v := range friends {
		persons[k] = v
	}
	persons["Anne"] = 18

	fmt.Println(friends, persons)

}
