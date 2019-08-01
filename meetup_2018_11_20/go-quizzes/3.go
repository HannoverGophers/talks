package main

import (
	"fmt"
)

// https://twitter.com/davecheney/status/1053419185680744448
func main() {
	var a []int
	b := a[:]
	fmt.Println(b == nil)
}

// true
// false
// panic on line 9
