package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const gr = 100

	var wg sync.WaitGroup
	wg.Add(gr * 2)

	var n int = 0

	// declare  a mutex value
	var m sync.Mutex

	// data race
	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)

			// 2. Block access to the variable

			m.Lock()
			n++
			m.Unlock()
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)

			m.Lock()
			n--
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("The final value of n: ", n)

}
