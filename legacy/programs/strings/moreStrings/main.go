package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var1, var2 := 'a', 'b'

	fmt.Printf("Type: %T, Value: %d \n", var1, var2)

	str := "taraÃ"
	fmt.Println(len(str))
	fmt.Println("Byte (not rune) at position 1: ", str[1])

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Println()
	for i := 0; i < len(str); {

		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c", r)
		i += size

	}
	fmt.Println()

	for _, r := range str {
		fmt.Printf("%c", r)
	}

	fmt.Println()
	// Lenght in Bytes and Runes

	s1 := "Golang"
	fmt.Println(len(s1))

	s2 := "straÆ"
	fmt.Println(len(s2))

	s3 := "ݶ"
	fmt.Println(len(s3))

	n := utf8.RuneCountInString(s2)
	m := utf8.RuneCountInString(s3)

	fmt.Println(n, m)

	// slicing a string

	sliced := "IÆݶ love how this plays out"
	fmt.Println(sliced[0:2])

	rs := []rune(sliced)
	fmt.Printf("%T \n", rs)

	fmt.Println(string(rs[0:2]))

	// string package

	testString := "This is JUST a test StriNg i WaNt To USE!"
	p := fmt.Println

	result := strings.Contains("I love Go Programming!", "love")

	fmt.Printf("%T \n", p)
	p(result)

	result = strings.ContainsAny("success", "xys")
	p(result)

	result = strings.ContainsRune("golang", 'g')

	noOfSubstring := strings.Count("cheesee", "ee")
	p(noOfSubstring)

	p(strings.ToLower(testString))
	p(strings.ToUpper(testString))

	// comparing strings

	p("go" == "go")
	p("Go" == "go")

	// lower all strings

	p(strings.ToLower("Go") == strings.ToLower("go"))

	// best solution

	p(strings.EqualFold("GO", "go"))

	// repeat

	mynewStr := strings.Repeat("abcd", 10)
	p(mynewStr)

	// replace

	myReplacedStr := strings.Replace("192.168.0.1", ".", ":", -1)

	p(myReplacedStr)

	myStrReplacedAll := strings.ReplaceAll("192.168.21.0", ".", ":")
	p(myStrReplacedAll)

	// split

	split := strings.Split("a,b,c", ",")
	fmt.Printf("%T \n", split)
	p(split)

	split = strings.Split("Go for Go!", "")
	fmt.Println(split)

	stringToJoin := []string{"I", "might", "grow", "to", "like", "this"}
	var joinedString = strings.Join(stringToJoin, " ")
	p(joinedString)

	// Field

	myOtherStr := " Orange Green \n Blue Yellow"
	fields := strings.Fields(myOtherStr)
	p(fields)

	// trimming

	unTrimmed := strings.TrimSpace("\t  Goodbye    Windows, Welcome Linux! \n")
	fmt.Printf("%q \n", unTrimmed)

	

}
