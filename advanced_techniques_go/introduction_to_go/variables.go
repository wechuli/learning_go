package main

import (
	"fmt"
	"math"
)

var Global int = 1234

const AnotherGlobal = -5678

func variables_play() {
	var j int
	i := Global + AnotherGlobal
	fmt.Println("Initial j value: ", j)
	j = Global
	k := math.Abs(float64(AnotherGlobal))
	fmt.Printf("Global=%d, i=%d, j=%d, k=%.2f. \n", Global, i, j, k)
}
