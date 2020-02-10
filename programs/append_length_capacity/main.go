package main

import "fmt"

func main() {
	var nums []int
	fmt.Printf("%#v \n", nums)

	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 1, 2)

	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 1, 3)

	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	letters := []string{"A", "B", "C", "D", "E", "F"}
	fmt.Printf("Length: %d, Capacity: %d \n", len(letters), cap(letters))

	letters = append(letters[0:1], "X", "Y")
	fmt.Printf("Length: %d, Capacity: %d \n", len(letters), cap(letters))
	// fmt.Println(letters[4]) this will result in an error

	fmt.Println(letters[:6])

}
