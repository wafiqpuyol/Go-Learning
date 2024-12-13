package main

import (
	"fmt"
	"time"
)

type Channel struct {
	ch1 chan string
	qCh chan struct{}
}

func (c *Channel) NewChannel() *Channel {
	return &Channel{
		ch1: make(chan string, 128),
		qCh: make(chan struct{}),
	}
}

func (c *Channel) LoadAllChannel() {
	//
channelLoop:
	for {
		select {
		case <-c.qCh:
			fmt.Println("Shutting down server")
			break channelLoop
		case msg := <-c.ch1:
			fmt.Println("Received message: ", msg)
		}
	}
}

func (c *Channel) QuitChannel() {
	close(c.qCh)
}

func main() {
	c := &Channel{}
	c = c.NewChannel()

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("Shutting down initialized")
		c.QuitChannel()
	}()

	go func() {
		for i := 1; i < 100; i++ {
			time.Sleep(time.Millisecond * 400)
			c.ch1 <- fmt.Sprintf("Message %d", i)
		}
	}()

	c.LoadAllChannel()
}
