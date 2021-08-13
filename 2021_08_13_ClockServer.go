package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept() // blocking until connection is coming
		if err != nil {
			log.Print(err) // connection error. i.e. connx aborted
			continue       // ready for next good connection
		}

		go handleConn(conn) // do actual work to handle connection

	}
}

func handleConn(c net.Conn) {
	defer c.Close() // like finally in C#
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
