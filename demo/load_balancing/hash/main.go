package main

import "fmt"

func main() {
	var devices = []string{
		"192.168.1.2:80",
		"192.168.1.3:80",
		"192.168.1.4:80",
	}

	for i := 0; i < 10; i++ {
		idx := i % len(devices)
		fmt.Println(devices[idx])
	}
}
