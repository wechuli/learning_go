package main

import "fmt"

func main() {
	var cities [4]string
	fmt.Println(cities)

	grades := [3]float64{78.43, 859}
	fmt.Println(grades)

	var taskDone = [...]bool{true, false, true, true}
	fmt.Println(taskDone)

	cities[0] = "Mombasa"
	cities[1] = "Nairobi"
	cities[2] = "Webuye"
	cities[3] = "Narok"

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	for key, value := range cities {
		fmt.Println(key, value)
	}

	nums := [...]int{30, -1, -6, 90, -6}
	evenPositive := 0

	for _, value := range nums {
		if value >= 0 && value%2 == 0 {
			evenPositive++
		}
	}

	fmt.Println("The number of even positive numbers is: ", evenPositive)

}
