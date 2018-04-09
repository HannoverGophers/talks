package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	// time measurement
	start := time.Now()

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// increment waitgroup counter
		go func(sec int) {

			fmt.Printf("waiting %d seconds...\n", sec)

			// do some fancy stuff
			time.Sleep(time.Second * time.Duration(sec))

			// telling waitgroup to decrement the counter
			wg.Done()

			fmt.Printf("done after %d seconds!\n", sec)

		}(i)
	}

	// wait until all work is done
	wg.Wait()

	fmt.Println("Elaspsed: ", time.Now().Sub(start))
}
