package main

import (
	"fmt"
	"sync"
	"time"
)

func callAfterParticularTimeByBlocking(sec time.Duration, wg *sync.WaitGroup) {
	time.Sleep(sec * time.Second) // this blocks the go routines
	fmt.Printf("Printing after %d seconds \n", sec)
	wg.Done()
}

func callAfterParticularTimeByNonBlocking(sec time.Duration, wg *sync.WaitGroup) {
	fmt.Println(time.Now())
	timer := time.NewTimer(sec * time.Second) // this does not blocks the go routines
	firedAt := <-timer.C
	fmt.Printf("Printing after %d seconds and value is %v \n", sec, firedAt)
	wg.Done()
}

func ExampleTimer() {
	var wg sync.WaitGroup
	fmt.Println("Prints the main block")

	fmt.Printf("Prints current time %v \n", time.Now())                   // prints current time and date
	fmt.Printf("Adds two hours %v \n", time.Now().Add(2*time.Hour))       // Adds two hours
	fmt.Printf("Subtracts two hours %v \n", time.Now().Add(-2*time.Hour)) // Subtracts two hours
	now := time.Now()
	past := time.Now().Add(-10 * time.Hour)
	fmt.Printf("Diff between two dates or time %v \n", now.Sub(past))
	wg.Add(3)

	go callAfterParticularTimeByNonBlocking(2, &wg)
	go callAfterParticularTimeByBlocking(2, &wg) // block the main thread and the prints the data
	go callAfterParticularTimeByBlocking(3, &wg) // block the main thread and the prints the data
	wg.Wait()

}
