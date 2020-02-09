package main

import "fmt"

func main() {
	var cities []string
	fmt.Println("cities is equal to nil: ", cities == nil)
	fmt.Printf("cities %T \n", cities)

	// we cannot assign toa nil slice
	// cities[0] = "London"

	numbers := []int{2, 3, 4, 5, 5}
	fmt.Println(numbers)

	nums := make([]int, 2)
	fmt.Printf("%#v \n", nums)

	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends)

	myFriend := friends[0]
	fmt.Println("My friend is ", myFriend)

	friends[0] = "Gabriel"
	fmt.Println("My best friend is ", friends[0])

	for index, value := range numbers {
		fmt.Printf("index: %v, value: %v \n", index, value)
	}

	var n []int
	n = numbers
	fmt.Println(n)

	// comparing slices
	var nCompare []int
	fmt.Println(nCompare == nil)

	m := []int{}
	fmt.Println(m == nil)

	a, b := []int{1, 2, 3}, []int{1, 2, 3}

	// a = nil

	var eq = false
	for index, valueA := range a {
		if valueA != b[index] {
			eq = false
			break
		}
		eq = true
	}

	fmt.Println(eq)

	// Coping and appending to slices

	numbers2 := []int{2, 3}

	numbers2 = append(numbers2, 10)
	fmt.Println(numbers2)

	numbers2 = append(numbers2, 20, 30)
	fmt.Println(numbers2)

	n2 := []int{100, 200}
	numbers = append(numbers, n2...)
	fmt.Println(numbers)

	// copying

	src := []int{10, 20, 30}
	// dst := make([]int, len(src))
	dst := make([]int, 2)

	nn := copy(dst, src)
	fmt.Println(src, dst, nn)

}
