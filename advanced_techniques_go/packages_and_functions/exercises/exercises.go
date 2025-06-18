package exercises

import (
	"fmt"
	"sort"
)

func ExersieOnePlay() {
	unsorted := [3]int{5, 4, 6}

	fmt.Println(sortInts(unsorted[0], unsorted[1], unsorted[2]))
	fmt.Println(sortInts2(unsorted[0], unsorted[1], unsorted[2]))

}

func sortInts(a, b, c int) (x, y, z int) {
	var first, second, third int
	numbers := []int{a, b, c}
	sort.Ints(numbers)
	first, second, third = numbers[0], numbers[1], numbers[2]

	x, y, z = first, second, third

	return x, y, z
}

func sortInts2(a, b, c int) (int, int, int) {
	numbers := []int{a, b, c}
	sort.Ints(numbers)
	return numbers[0], numbers[1], numbers[2]
}
