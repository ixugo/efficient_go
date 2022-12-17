package main

// broken pipe

import (
	"errors"
	"log"
	"net"
	"os"
	"syscall"
	"time"
)

func server() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("server", err)
		os.Exit(1)
	}
	data := make([]byte, 1)
	if _, err := conn.Read(data); err != nil {
		log.Fatal("server", err)
	}

	conn.Close()
}

func client() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("client", err)
	}

	// write to make the connection closed on the server side
	if _, err := conn.Write([]byte("a")); err != nil {
		log.Printf("client: %v", err)
	}

	time.Sleep(1 * time.Second)

	// write to generate an RST packet
	if _, err := conn.Write([]byte("b")); err != nil {
		log.Printf("client: %v", err)
	}

	time.Sleep(1 * time.Second)

	// write to generate the broken pipe error
	if _, err := conn.Write([]byte("c")); err != nil {
		log.Printf("client: %v", err)
		if errors.Is(err, syscall.EPIPE) {
			log.Print("This is broken pipe error")
		}
	}
}

func main() {
	go server()

	time.Sleep(3 * time.Second) // wait for server to run

	client()
}
