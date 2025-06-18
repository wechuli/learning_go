package main

import (
	"fmt"
	"os"
	"strconv"
)

func DeleteSlice() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need an integer value")
		return
	}
	index := arguments[1]
	i, err := strconv.Atoi(index)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Using index", i)

	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice: ", aSlice)

	// Delete an elelment at index i

	if i > len(aSlice)-1 {
		fmt.Println("Cannot delete element", i)
		return
	}

	fmt.Println("Slice: aSlice[:2]", aSlice[:2])

	aSlice = append(aSlice[:i], aSlice[i+1:]...)

	fmt.Println("After 1st deletion: ", aSlice)

}
