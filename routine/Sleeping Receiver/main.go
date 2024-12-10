package main

import (
	"fmt"
	"time"
)

func a(messages chan string) {
	fmt.Println("2")
	msg := <-messages
	// time.Sleep(time.Second * 1) // try to toggle this sleep see what happens
	fmt.Println("3")
	fmt.Println(msg)
}
func main() {
	messages := make(chan string)
	// Receiver
	go a(messages)
	// Sender
	fmt.Println("1")
	time.Sleep(2000 * time.Millisecond)
	messages <- "Message : Hello World"
	fmt.Println("4")

	// Wait spawned goroutine process
	time.Sleep(1000 * time.Millisecond)
	close(messages)
}
