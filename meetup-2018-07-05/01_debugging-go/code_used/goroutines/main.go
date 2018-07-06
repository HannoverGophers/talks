package main

import "fmt"

func f(ch chan int) {
	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	c := make(chan int)

	go f(c)

	for i := range c {
		fmt.Printf("%d\n", i)
	}

}
