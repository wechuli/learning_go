package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func(value float64, wg *sync.WaitGroup) {
		fmt.Printf("%.2f \n", math.Sqrt(value))
		wg.Done()
	}(85.6, &wg)

	for i := 100; i <= 149; i++ {
		wg.Add(1)

		go func(value float64, wg *sync.WaitGroup) {
			fmt.Printf("%.2f \n", math.Sqrt(value))
			wg.Done()
		}(float64(i), &wg)

	}

	wg.Wait()
}
