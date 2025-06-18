package main

import (
	"fmt"
	"sync"
)

func deposit(b *int, n int, wg *sync.WaitGroup, mu *sync.Mutex) {

	mu.Lock()
	*b += n
	mu.Unlock()
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, mu *sync.Mutex) {

	mu.Lock()
	*b -= n
	mu.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	// declare a mutex value
	var m sync.Mutex

	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &m)
		go withdraw(&balance, i, &wg, &m)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}
