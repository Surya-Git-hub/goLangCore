package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in golang")

	myCh := make(chan int, 1)
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// close(myCh)

		value, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(value)
		fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		defer wg.Done()
	}(myCh, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 5
		// myCh <- 6
		defer wg.Done()
	}(myCh, wg)

	wg.Wait()
}
