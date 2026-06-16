package main

import (
	"fmt"
)

// https://twitter.com/davecheney/status/1039997464361623552
func main() {
	fmt.Println(string('7' << 1))
}

// 7
// 14
// n
// doesn't compile
