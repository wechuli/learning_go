package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	// a[start:stop]
	b := a[1:3]
	fmt.Printf("%v, %T \n", b, b)

	var s1 = []int{1, 2, 3, 4, 5, 6, 7}
	s2 := s1[1:4]
	fmt.Println(s2)
	fmt.Println(s1[2:])
	fmt.Println(s1[:len(s1)])
	fmt.Println(s1[:])

	// append with slicing

	s1 = append(s1[:4], 100)
	fmt.Println(s1)

	// backing array
	sArray := []int{10, 20, 30, 40, 50}
	s3, s4 := sArray[0:2], sArray[1:3]

	s3[1] = 600 // this modifies the backing array for sArray

	fmt.Println(sArray)
	fmt.Println(s4)

	arr1 := [4]int{10, 20, 30, 40}

	slice1, slice2 := arr1[0:2], arr1[1:3]

	arr1[1] = 2
	slice2[0] = 545454 // modifies the array

	fmt.Println(arr1, slice1, slice2)

	// use append to  create a new cooy of the slice

	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)
	cars[0] = "Nissan"
	fmt.Println(cars, newCars)

	anotherSlice := []int{10, 20, 30, 40, 50}
	newSlice := anotherSlice[0:3]

	fmt.Println(len(newSlice), cap(newSlice))

}
