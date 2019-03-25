package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net"
	"time"
)

func StartServer(server SSHServer, lines []string) {
	log.Infof("Starting ssh tar pit on %v:%v", server.Host, server.Port)
	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", server.Port))

	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	for {
		log.Info("Starting to accept new connections")
		conn, err := ln.Accept()
		handleError(err)

		if len(lines) == 0 {
			go HandleConnectionNumber(conn)
		} else {
			go HandleConnectionText(conn, lines)
		}
	}
}

type SSHServer struct {
	Host string
	Port int
}

func HandleConnectionNumber(conn net.Conn) {
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

func HandleConnectionText(conn net.Conn, lines []string) {
	log.Infof("Connection opened for %v", conn.RemoteAddr())

	lineNumber := 0
	for {
		output := fmt.Sprintf("%v\r\n", lines[lineNumber])
		log.Debugf("Wasting 10 seconds: %v", output)
		_, err := conn.Write([]byte(output))
		if err != nil {
			log.Infof("Connection closed for %v", conn.RemoteAddr())
			break
		}
		time.Sleep(10 * time.Second)

		lineNumber = lineNumber + 1
		if lineNumber > len(lines)-1 {
			lineNumber = 0
		}
	}
	log.Infof("Connection opened for %v", conn.RemoteAddr())
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
