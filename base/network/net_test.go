package network

import (
	"context"
	"fmt"
	"io"
	"net"
	"strings"
	"testing"
	"time"
)

// Server function to listen on UDP and print the received data length.
func server() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":8080")
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
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}
		total += n
		fmt.Printf("Received %d total: %d\n", n, total)
	}
}

const blockSize = 1500

func client() {
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
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

func TestASD(t *testing.T) {
	go server()
	go client()
	// Let the program run for a while so the server can receive the data.

	time.Sleep(10 * time.Second)
}
func echoServerUDP(ctx context.Context, addr string) (net.Addr, error) {

	s, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.ParseIP(addr), Port: 0})
	if err != nil {
		return nil, fmt.Errorf("binding to udp %s: %w", addr, err)
	}

	go func() {
		go func() {
			<-ctx.Done()
			_ = s.Close()
		}()

		buf := make([]byte, 1024)
		for {
			n, cli, err := s.ReadFrom(buf) // client to server
			if err != nil {
				return
			}

			_, err = s.WriteTo(buf[:n], cli) // server to client
			if err != nil {
				return
			}
			// fmt.Println("server>>>", string(buf[:n]))
		}
	}()

	return s.LocalAddr(), nil
}

func TestEchoServerUDP(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	serverAddr, err := echoServerUDP(ctx, "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	defer cancel()

	cli, err := net.Dial("udp4", serverAddr.String())
	if err != nil {
		t.Fatal(err)
	}
	client := cli.(*net.UDPConn)

	// client.SetWriteBuffer(5)
	// client.SetReadBuffer(5)

	go func() {
		for {
			buf := make([]byte, 1500)
			n, _, err := client.ReadFrom(buf)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println(len(buf[:n]))
		}
	}()
	var a string
	for range 1200 {
		a += "abc"
	}
	msg := []byte(a)
	for i := 0; i < len(msg); i++ {
		i += 1024
		_, err = client.Write(msg[:1024])
		if err != nil {
			fmt.Println(err)
			return
		}
		msg = msg[1024:]
	}
	time.Sleep(5 * time.Second)

}
func TestABC(t *testing.T) {
	addr, _ := echoServerUDP(context.Background(), ":6234")

	d, err := net.Dial("udp4", addr.String())
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		buf := make([]byte, 5)
		for {
			n, err := io.ReadFull(d, buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(buf[:n]))
		}
	}()

	d.Write([]byte("123"))
	// time.Sleep(time.Second)
	if _, err := d.Write([]byte("45678915111241241")); err != nil {
		t.Fatal(err)
	}
	d.Write([]byte("abc"))
	// d.Write([]byte("45678900"))

	time.Sleep(5 * time.Second)

}
