package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var ladies = [3]string{"June", "Jess", "Mercy"}

	for i := 0; i < len(ladies); i++ {
		println(ladies[i])
	}

	// Mimicking a while loop

	// for i := 10; i > 0; {
	// 	fmt.Println("I will print forever and ever")
	// }
	// // or

	// for {
	// 	fmt.Println("I am an unreachable code")
	// }

	// continue
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	// break
	count := 0
	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13 \n", i)
			count++
		}
		if count == 10 {
			break
		}
	}
	fmt.Println(count)
}
