package main

import "fmt"

func TypeAssertionPlay() {
	anInt := returnNumber()
	number := anInt.(int)
	number++
	fmt.Println(number)

	value, ok := anInt.(int64)
	if ok {
		fmt.Println("Type assertion successful: ", value)
	} else {
		fmt.Println("Type assertion failed!")
	}

	_ = anInt.(bool)
}
func returnNumber() interface{} {
	return 12
}
