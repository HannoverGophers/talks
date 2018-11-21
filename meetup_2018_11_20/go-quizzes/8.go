package main

import "fmt"

type Q struct{}

func (Q) Hola() string { return "Bueno" }

// https://twitter.com/davecheney/status/1031389514890006528
func main() {
	var q Q
	fmt.Println(Q.Hola(q))
}

// Bueno
// Good
// Nada, it won't compile
