package main

import (
	"fmt"
	"os"
	"regexp"
)

func RegexPlay() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> string.")
		return
	}

	s := arguments[1]
	ret := matchName(s)

	fmt.Println("Match name: ", ret)

	fmt.Println("Match digit: ", matchInts(s))

}

func matchName(s string) bool {
	t := []byte(s)

	re := regexp.MustCompile(`^[A-Z][a-z]*$`)

	return re.Match(t)

}

func matchInts(s string) bool {

	re := regexp.MustCompile(`^[-+]?\d+$`)

	return re.MatchString(s)

}
