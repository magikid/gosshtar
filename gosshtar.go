package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	log.Printf("Connection opened for %v", conn.RemoteAddr())
	for {
		output := fmt.Sprintf("%v\r\n", rand.Int())
		log.Print("Wasting 10 seconds: ", output)
		_, err := conn.Write([]byte(output))
		if err != nil {
			log.Printf("Connection closed for %v", conn.RemoteAddr())
			break
		}
		time.Sleep(10 * time.Second)
	}
}

func main() {
	log.Print("Starting ssh tar pit")
	ln, err := net.Listen("tcp", ":2222")

	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	for {
		log.Print("Starting to accept new connections")
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}
