package ip

import (
	"fmt"
	"testing"
)

func TestLocalIP(t *testing.T) {
	ip := LocalIP()
	fmt.Println(ip)
}

func TestExternalIP(t *testing.T) {
	ip, err := ExternalIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
}
