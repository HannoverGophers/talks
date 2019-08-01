package main

import (
	"fmt"
	"runtime"
	"sync"
)

// https://twitter.com/davecheney/status/1031699003803463680
func main() {
	var wg sync.Waitgroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		runtime.Goexit()
	}()
	wg.Wait()
	fmt.Println("Hello, playground")
}

// Hello, playground
// It panics
// Nothing, it don'st compile
