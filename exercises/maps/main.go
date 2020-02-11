package main

import "fmt"

func main() {

	// Question 1

	var m2344 map[string]float64

	fmt.Println(m2344 == nil)

	m2 := map[int]string{
		1: "One",
		2: "Two",
	}
	fmt.Println(m2)

	m2[10] = "Abba"

	fmt.Println(m2)

	k, ok := m2[58656]
	fmt.Println(k, ok)

	k2, ok2 := m2[1]
	fmt.Println(k2, ok2)

	// Question 2
	m1 := map[int]bool{}
	m1[5] = true

	m3 := map[int]int{3: 10, 4: 40}
	m4 := map[int]int{4: 40, 3: 10}

	// fmt.Println(a == b) //invalid
	aString := fmt.Sprintf("%s", m3)
	bString := fmt.Sprintf("%s", m4)

	//fmt.Println(m2 == m3)

	fmt.Println(aString == bString)
}
