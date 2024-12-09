package main

import (
	"fmt"
)

func receiver(emailChan chan string, receiverChan chan bool) {
	defer func() { receiverChan <- true }()

	for v := range emailChan {
		fmt.Println("Received email:", v)
		// time.Sleep(time.Second * 1)
	}
}

func main() {
	emailChan := make(chan string)
	receiverChan := make(chan bool)

	go receiver(emailChan, receiverChan)
	for i := 0; i < 10; i++ {
		emailChan <- fmt.Sprintf("email %d", i)
	}
	fmt.Println("done Sending ")
	// close(emailChan)
	fmt.Println(<-receiverChan)
}
