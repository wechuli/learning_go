package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 255
	x++
	fmt.Println(x) // the overflow is silent

	// the error may be caught if at compile tume

	// var anotherX int8 = 256

	// fmt.Println(anotherX)

	// float overflow

	f := float32(math.MaxFloat32)
	f = f * 1.2 //f overflows to infinet

	fmt.Println(f)

}
