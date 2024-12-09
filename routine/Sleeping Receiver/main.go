package main

import (
	"fmt"
	"time"
)

func a(messages chan string) {
	fmt.Println("Receiver : I am waiting for your message.")
	msg := <-messages // will block this execution
	fmt.Println("Receiver : I got a mail.")
	fmt.Println(msg)
}
func main() {
	messages := make(chan string)
	// Receiver
	go a(messages)
	// Sender
	time.Sleep(2000 * time.Millisecond)
	messages <- "Message : Do you like go language?"
	// Wait spawned goroutine process
	time.Sleep(1000 * time.Millisecond)
}
