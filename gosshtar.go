package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net"
	"os"
	"time"
)

func handleConnection(conn net.Conn) {
	log.Infof("Connection opened for %v", conn.RemoteAddr())
	for {
		output := fmt.Sprintf("%v\r\n", rand.Int())
		log.Debugf("Wasting 10 seconds: %v", output)
		_, err := conn.Write([]byte(output))
		if err != nil {
			log.Infof("Connection closed for %v", conn.RemoteAddr())
			break
		}
		time.Sleep(10 * time.Second)
	}
}

func main() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)

	debug := os.Getenv("DEBUG")
	if debug != "" {
		log.SetLevel(log.DebugLevel)
	}

	log.Info("Starting ssh tar pit")
	ln, err := net.Listen("tcp", ":2222")

	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	for {
		log.Info("Starting to accept new connections")
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}
