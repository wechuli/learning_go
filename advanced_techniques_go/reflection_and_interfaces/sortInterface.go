package main

import (
	"fmt"
	"sort"
)

type S1 struct {
	F1 int
	F2 string
	F3 int
}

type S2 struct {
	F1 int
	F2 string
	F3 S1
}

type S2Slice []S2

func SortInterfacePlay() {
	data := []S2{
		{1, "One", S1{1, "S1_1", 10}},
		{2, "Two", S1{2, "S1_1", 20}},
		{-1, "Two", S1{-1, "S1_1", -20}},
	}

	fmt.Println("Before: ", data)
	sort.Sort(S2Slice(data))
	fmt.Println("After: ", data)

	// Reverse sorting works automatially

	sort.Sort(sort.Reverse(S2Slice(data)))
	fmt.Println("Reverse: ", data)
}

func (a S2Slice) Len() int {
	return len(a)
}

func (a S2Slice) Less(i, j int) bool {
	return a[i].F3.F1 < a[j].F3.F1
}

func (a S2Slice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
