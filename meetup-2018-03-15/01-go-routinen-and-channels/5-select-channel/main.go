package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)  // returns receiving-only channel
	boom := time.After(5 * time.Second) // returns receiving-only channel

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
