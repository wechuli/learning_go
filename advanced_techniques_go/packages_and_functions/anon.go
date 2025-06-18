package main

import "fmt"

func AnonymousFuncPlay() {
	val := 10
	anF := func(param int) int {
		return param * param
	}
	fmt.Println(anF(val))
}
