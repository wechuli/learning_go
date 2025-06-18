package main

import "fmt"

func ArraySlicePlay() {
	aSlice := []float64{}

	// Both length and capacity are 0 because aSlice is empty

	fmt.Println(aSlice, len(aSlice), cap(aSlice))

	// Add elements to a slice

	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -34.0)

	fmt.Println(aSlice, "with length", len(aSlice), "and capacity", cap(aSlice))

	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4

	// Now we will need to use append

	t = append(t, -5)
	fmt.Println(t)

	//two D slice

	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}

	for _, i := range twoD {
		for _, k := range i {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	make2D := make([][]int, 2)

	fmt.Println(make2D)

	make2D[0] = []int{1, 2, 3, 4}
	make2D[1] = []int{-1, -2, -3, -4}
	fmt.Println(make2D)

}

func CapLengthPlay() {
	// Only length is defined. Capacity = length

	a := make([]int, 4)
	fmt.Println("L: ", len(a), "C: ", cap(a))

	// Initialize slice. Capacity = length

	b := []int{0, 1, 2, 3, 4}
	fmt.Println("L:", len(b), "C:", cap(b))

	// Same length and capacity
	aSlice := make([]int, 4, 4)
	fmt.Println(aSlice)
	// Add an element
	aSlice = append(aSlice, 5)
	fmt.Println(aSlice)
	// The capacity is doubled
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))
	// Now add four elements
	aSlice = append(aSlice, []int{-1, -2, -3, -4}...)
	fmt.Println(aSlice)
	// The capacity is doubled
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))
}

func SliceSectionsPlay() {
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(aSlice)
	l := len(aSlice)

	// First 5 elements
	fmt.Println(aSlice[0:5])
	// First 5 elements
	fmt.Println(aSlice[:5])

	// Last 2 elements
	fmt.Println(aSlice[l-2 : l])

	// Last 2 elements
	fmt.Println(aSlice[l-2:])

	// First 5 elements
	t := aSlice[0:5:10]
	fmt.Println(len(t), cap(t))

	// Elements at indexes 2,3,4
	// Capacity will be 10-2
	t = aSlice[2:5:10]
	fmt.Println(len(t), cap(t))

	// Elements at indexes 0,1,2,3,4
	// New capacity will be 6-0
	t = aSlice[:5:6]
	fmt.Println(len(t), cap(t))
}
