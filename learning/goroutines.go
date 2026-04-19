package main

import (
	"fmt"
	"sync"
)

func ExampleGoRoutines() {
	var counter int
	// var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// mu.Lock() // Locking the counter
			counter++
			// mu.Unlock() // Unlocking after update
		}()
	}

	wg.Wait() // wait for all goroutines to finish.
	fmt.Println("Final Counter:", counter)

}
