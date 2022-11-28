package ip

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestLocalIP(t *testing.T) {
	conn, err := net.DialTimeout("udp", "8.8.8.8:53", 3*time.Second)
	if err != nil {
		t.Fatal(err)

	}
	host, _, _ := net.SplitHostPort(conn.LocalAddr().(*net.UDPAddr).String())
	fmt.Println(host)
}

func TestExternalIP(t *testing.T) {
	ip, err := ExternalIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
}
