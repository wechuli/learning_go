package problems

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type Animal struct {
	Name  string
	Sound string
}

type persons []Person

func Problem1Play() {
	people := []Person{
		{"Alice", 25},

		{"Dave", 35},
		{"Bob", 30},
		{"Alice", 12},
		{"Charlie", 20},
	}

	persons := persons(people)
	sort.Sort(persons)

	fmt.Println(persons)

	p := Person{"Alice", 25}
	a := Animal{"Dog", "Woof"}

	var fraction float64 = 45.3

	describe(p)
	describe(a)
	describe(fraction)

}

func (p persons) Len() int {
	return len(p)
}

func (p persons) Less(i, j int) bool {
	if p[i].Name == p[j].Name {
		return p[i].Age < p[j].Age
	}
	return p[i].Name < p[j].Name
}

func (p persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func describe(i interface{}) {
	switch v := i.(type) {
	case Animal:
		fmt.Printf("This is an animal called %s, which maked the sound '%s' \n", v.Name, v.Sound)
	case Person:
		fmt.Printf("This is a person named %s, who is %d years old \n", v.Name, v.Age)
	default:
		fmt.Println("Unknown type")

	}
}
