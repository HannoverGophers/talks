package main

import (
	"fmt"
	"time"
)

func main() {

	ping := make(chan bool)
	pong := make(chan bool)
	defer Close(ping)
	defer Close(pong)

	// Ping listen to the ping-channel and send a message to the pong-channel
	go Ping(ping, pong)
	// Pong listen to the pong-channel and send a message to the ping-channel
	go Pong(pong, ping)

	// initial start
	ping <- true

	time.Sleep(time.Second * 10)
}

func Ping(listen chan bool, send chan bool) {
	for {
		<-listen
		fmt.Println("Ping!")
		time.Sleep(time.Millisecond * 500)
		send <- true
	}
}

func Pong(listen chan bool, send chan bool) {
	for {
		<-listen
		fmt.Println("Pong!")
		time.Sleep(time.Millisecond * 500)
		send <- true
	}
}

func Close(c chan bool) {
	fmt.Println("Closing channel...")
	close(c)
}
