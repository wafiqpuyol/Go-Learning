package main

import (
	"fmt"
	"time"
)

func msgReceiveAndSendBack(ch chan string, msg string) {
	fmt.Println("Message received: ", <-ch) //blocking code
	time.Sleep(3 * time.Second)
	ch <- msg //blocking code
}

func main() {
	ch := make(chan string)
	fmt.Println("Channel created")          // synchronous code
	go msgReceiveAndSendBack(ch, "Hello")   // synchronous code
	ch <- "hello world"                     //blocking code
	fmt.Println("Receiving Message")        // synchronous code
	fmt.Println("Message received: ", <-ch) //blocking code
}
