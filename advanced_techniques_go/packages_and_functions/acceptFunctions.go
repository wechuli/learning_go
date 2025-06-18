package main

import (
	"fmt"
	"math"
	"sort"
)

type Grades struct {
	Name    string
	Surname string
	Grade   int
}

func customOp(value float64) float64 {
	return value + 3
}
func apply(f func(float64) float64, x float64) float64 {
	return f(x)
}
func AcceptFuncPlay() {
	fmt.Println(apply(math.Sin, math.Pi/2))
	fmt.Println(apply(math.Cos, math.Pi))
	fmt.Println(apply(math.Tan, math.Pi/4))
	fmt.Println(apply(customOp, 45.34))

	data := []Grades{{"J.", "Lewis", 10}, {"M.", "Tsoukalos", 7},
		{"D.", "Tsoukalos", 8}, {"J.", "Lewis", 9}}

	isSorted := sort.SliceIsSorted(data, func(i, j int) bool {
		return data[i].Grade < data[j].Grade
	})

	if isSorted {
		fmt.Println("It is sorted!")
	} else {
		fmt.Println("It is NOT sorted!")
	}

	sort.Slice(data,
		func(i, j int) bool { return data[i].Grade < data[j].Grade })
	fmt.Println("By Grade:", data)
}
