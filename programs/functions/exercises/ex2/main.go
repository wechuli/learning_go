package main

import (
	"fmt"
	"strconv"
	"strings"
)

func myFunc(strNumber rune) int64 {

	valueToString := string(strNumber)

	// result, _ := strconv.ParseInt(string(strNumber), 10, 64)

	var mult int64 = 0

	for i := 1; i <= 10; i++ {
		result, _ := strconv.ParseInt(strings.Repeat(valueToString, i), 10, 64)
		mult += result
	}
	return mult
}
func main() {
	fmt.Println(myFunc('9'))
	fmt.Println(myFunc('5'))
}
