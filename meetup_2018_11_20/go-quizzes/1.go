package main

import (
	"fmt"
)

// https://twitter.com/davecheney/status/1053490060492890118
func main() {
	one, two, three := 0.1, 0.2, 0.3
	fmt.Println(one+two > three)
}

// True
// False
// 0.3
