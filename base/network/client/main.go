package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	client()
}

const blockSize = 1500

func client() {
	udpAddr, err := net.ResolveUDPAddr("udp", "39.101.79.234:38080")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Failed to connect to UDP server:", err)
		return
	}
	defer conn.Close()

	message := strings.Repeat("A", 10000) // Sending 10000 'A's
	for i := 0; i < len(message); i += blockSize {
		end := i + blockSize
		if end > len(message) {
			end = len(message)
		}
		_, err = conn.Write([]byte(message[i:end]))
		if err != nil {
			fmt.Println("Error sending data:", err)
			break
		}
	}
}
