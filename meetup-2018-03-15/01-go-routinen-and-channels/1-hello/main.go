package main

import (
	"fmt"
	"time"
)

func hello(name string) {
	time.Sleep(time.Second * 2)
	fmt.Printf("Hello %s from go routine!\n", name)
}

func main() {

	fmt.Println("Hello from main()")

	go hello("Gopher")
	go hello("John")

	fmt.Println("Waiting ...")
	time.Sleep(time.Second * 5)

	fmt.Println("Bye!")
}
