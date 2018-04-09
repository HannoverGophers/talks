package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan struct{})

	// starting anonymous function
	go func() {

		fmt.Println("waiting for some time...")

		// do some fancy stuff
		time.Sleep(time.Second * 2)

		done <- struct{}{}
		fmt.Println("end of go routine")
	}()

	// blocking until receiving message
	<-done

	<-time.After(time.Second * 1)

	fmt.Println("done!")
}
