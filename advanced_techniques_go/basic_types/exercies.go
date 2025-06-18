package main

import (
	"fmt"
)

func ExercisePlay() {

	a1 := [3]int{1, 2, 3}
	a2 := [3]int{3, 4, 5}

	fmt.Println(concatenate(a1, a2))

	//concatenate2

	arr1 := [3]int{1, 2, 3}
    arr2 := [3]int{4, 5, 6}
    newArr := concatenate2(arr1, arr2)
    fmt.Println(newArr)

	//concatenate slices

	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	newArray := concatenateSlices(slice1, slice2)
	fmt.Println(newArray)
}

func concatenate(a1 [3]int, a2 [3]int) []int {
	var s []int
	s = a1[:]
	s = append(s, a2[:]...)

	return s
}

func concatenate2(a [3]int, b [3]int) [6]int {
	var c [6]int
	var fullSlice = append(a[:],b[:]...)

	for i,item := range fullSlice{
		c[i] = item
	}

	return c
}

func concatenateSlices(slice1 []int, slice2 []int) [6]int {
	var newArray [6]int

	fullSlice := append(slice1, slice2...)

	for i, item := range fullSlice {
		newArray[i] = item
	}

	return newArray
}
