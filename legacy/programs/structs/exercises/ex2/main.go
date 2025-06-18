package main

import "fmt"

type grade struct {
	course string
	marks  float32
}
type person struct {
	name      string
	age       int
	favColors []string
	grades    grade
}

func main() {
	me := person{
		name:      "Paul",
		age:       27,
		favColors: []string{"Teal", "Maroon", "DarkBlue", "LightPink"},
		grades: grade{course: "Math",
			marks: 84}}

	fmt.Printf("%#+v \n", me)

	me.grades.course = "Golang"
	me.grades.marks = 98

	fmt.Printf("%#+v \n", me)

}
