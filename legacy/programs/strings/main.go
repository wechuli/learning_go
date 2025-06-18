package main

import "fmt"

func main() {
	myString := "Hi there go !"
	fmt.Println(myString)

	// raw strings
	s2 := `This is a raw string \n , "we can use this"`
	s3 := `C:\Users\Wechuli`
	fmt.Println(s2, s3)

	// Concatentatre

	var s4 = "I love " + "Go " + "Programming"
	fmt.Println(s4)

	fmt.Println(s4[0])
	// strings are immutable

	// Runes

	var1, var2 := 'a', 'b'
	fmt.Printf("Type: %T, Value: %d \n", var1, var2)

	str := "þÿ[Repe"

	fmt.Println(str)
	fmt.Println(str[0])
	fmt.Printf("%T \n", str[0])
}
