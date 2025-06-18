package main

import (
	"fmt"
	"math"
	"time"
)

const secondsInHour = 3600

func main() {
	fmt.Println(" Hello world")

	// var leftHand float64 = 89.09
	distance := 60.8
	fmt.Println(time.Now())
	fmt.Println(math.Abs(-333))

	fmt.Printf("The distance in miles is %f \n", distance*0.685)
	fakeFunction()
}

func fakeFunction() {
	fmt.Println("hi there")
}
