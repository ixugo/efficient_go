package main

import (
	"fmt"
	"net"
)

// Server function to listen on UDP and print the received data length.
func server() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":38080")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to start UDP server:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024) // A buffer to read data
	total := 0
	for {
		n, raddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}
		total += n
		fmt.Printf("raddr[%s] Received %d total: %d\n", raddr.String(), n, total)
	}
}

const blockSize = 1500

func main() {
	server()
}
