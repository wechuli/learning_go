package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// Question 1
	var name string = "Paul Wechuli"
	country := "Kenya"

	fmt.Printf("Your name: %s \n", name)
	fmt.Printf("Country: %s \n", country)

	fmt.Printf(`He says: "Hello"`)
	fmt.Println()
	fmt.Printf(`C:\Users\a.txt`)
	fmt.Println()
	// Question 2

	r := 'ă'

	fmt.Printf("Type of r: %T \n", r)

	str1, str2 := "ma", "m"

	newString := str1 + str2 + string(r)
	fmt.Println(newString)

	// Question 3

	s1 := "țară means country in Romanian"

	for i := 0; i < len(s1); i++ {
		fmt.Printf("%d \n", s1[i])
	}

	for i := 0; i < len(s1); {
		result, size := utf8.DecodeRuneInString(s1[i:])

		fmt.Printf("%c ", result)
		fmt.Println(result, size)
		i += size

	}

	fmt.Println(s1[2:])

	// Question 5
	s := "你好 Go!"

	// converting string to byte slice
	b := []byte(s)

	// printing out the byte slice
	fmt.Printf("%#v \n", b)

	// iterating over the byte slice and printing out each index and byte in the byte slice

	for i, v := range b {
		fmt.Printf("%d -> %d\n", i, v)
	}

	runeSlice := []rune(s)

	// printing out the rune slice
	fmt.Printf("%#v \n", runeSlice)

	// iterating over the rune slice and printing out each index and rune in the rune slice

	for i, v := range runeSlice {
		fmt.Printf("%d -> %d -> %c\n", i, v, v)
	}

}
