package main

import "fmt"

func LoopThroughString(value string) {
	for _, sval := range value {
		fmt.Println(string(sval))
	}
}
