package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func loop(conn net.Conn) {
	buffer := make([]byte, 1024)

	if _, err := conn.Read(buffer); err != nil {
		log.Fatal("Read error: ", err)
	}

	time.Sleep(time.Second * 10)

	conn.Write([]byte("HTTP/1.1 200 MUTH\r\n\r\n"))
	conn.Close()

	log.Println("Closed connection")
}
func main() {

	listener, err := net.Listen("tcp", "localhost:4550")
	fmt.Println(listener.Addr())
	if err != nil {
		log.Fatal("Listen error: ", err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}

		go loop(conn)
		log.Println("Accepted connection")
	}
}
