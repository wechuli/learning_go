package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Record struct {
	firstName string
	lastName  string
	tel       string
}

func MatchRecordPlay() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> record")
		return
	}

	s := arguments[1]
	err := matchRecord(s)
	fmt.Println(err)
}

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)

	return re.Match(t)
}

func matchTel(s string) bool {
	re := regexp.MustCompile(`\d+$`)
	return re.MatchString(s)
}

func matchRecord(s string) bool {
	fields := strings.Split(s, ",")

	if len(fields) != 3 {
		return false
	}

	record := Record{firstName: fields[0], lastName: fields[1], tel: fields[2]}

	if !matchNameSur(record.firstName) {
		return false
	}
	if !matchNameSur(record.lastName) {
		return false
	}

	return matchTel(record.tel)
}
