package main

import "fmt"

type Person struct {
	name string
	age  int
}

func MapPlay() {
	// basicsPlay()
	//nilPlay()
	mapIteration()
}

func basicsPlay() {
	m := map[string]int{
		"key1": -1,
		"key2": 123}

	p := Person{"Paul", 32}

	m2 := map[string]Person{
		"one": p}

	fmt.Println(m)

	v, ok := m2["fake"]

	if ok == false {
		fmt.Println(v)
	}

	fmt.Println("Length of the map: ", len(m))

	nilPlay()

}

func nilPlay() {
	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Println("aMap:", aMap)
	aMap = nil
	fmt.Println("aMap:", aMap)
	if aMap == nil {
		fmt.Println("nil map!")
		aMap = map[string]int{}
	}
	aMap["test"] = 1

	// This will crash!
	// aMap = nil
	// aMap["test"] = 1
}

func mapIteration() {
	aMap := make(map[string]string)
	aMap["123"] = "456"
	aMap["key"] = "A value"

	for key, v := range aMap {
		fmt.Println("Key:", key, "value", v)
	}
}
