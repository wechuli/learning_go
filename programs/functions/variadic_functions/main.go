package main

import (
	"fmt"
	"strings"
)

func f1(a ...int) {
	fmt.Printf("%T \n", a)
	fmt.Printf("%#v \n", a)
}

func f2(a ...int) {
	a[0] = 50
}

func sumAndProduct(values ...float64) (sum float64, product float64) {
	sum = 0
	product = 1
	for _, value := range values {
		sum += value
		product *= value
	}

	return sum, product

}

// mixing variadic parameters
func personInformation(age int, names ...string) string {
	fullName := strings.Join(names, " ")
	returnString := fmt.Sprintf("Age: %d, Full Name: %s", age, fullName)

	return returnString
}
func main() {

	f1(45, 45, 34, 5)
	f1()

	nums := []int{1, 2}
	nums = append(nums, 3, 4, 5)
	fmt.Println(nums)

	f1(nums...)

	f2(nums...)

	fmt.Println("Nums: ", nums)

	s, p := sumAndProduct(2, 34.6, 645, 3, 1, 85, 8)
	fmt.Println(s, p)

	// mixing

	info := personInformation(27, "Wolfgang", "Amadeus", "Mozart")
	fmt.Println(info)

}
