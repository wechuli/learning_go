package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func sum(first float64, second float64, wg *sync.WaitGroup) {
	result := first + second
	fmt.Printf("Sum: %.2f \n", result)
	wg.Done()
}
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		first := rand.NormFloat64()
		second := rand.NormFloat64()

		println(first, second)

		sum(first, second, &wg)
	}

	wg.Wait()
}
