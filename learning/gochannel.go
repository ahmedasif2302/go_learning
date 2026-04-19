package main

import (
	"fmt"
	"sync"
)

func ExampleOfChannels() {

	// var wg sync.WaitGroup
	wg := &sync.WaitGroup{} // address is passed here
	myCh := make(chan int, 2)

	wg.Add(2)
	// ch chan<- int <- meaning set value in the channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 10
		// ch <- 11
		defer wg.Done()
		close(ch)

	}(myCh, wg) // we can pass direct var of the address

	// ch <-chan int <- meaning read value in the channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch
		if !isChannelOpen {
			fmt.Println(val, isChannelOpen)
			return
		}
		fmt.Println(val, isChannelOpen)
		defer wg.Done()

	}(myCh, wg)

	wg.Wait()
}
