package main

import "fmt"

var devices = []string{
	"192.168.1.2:80",
	"192.168.1.3:80",
	"192.168.1.4:80",
	"192.168.1.5:80",
	"192.168.1.6:80",
	"192.168.1.7:80",
	"192.168.1.8:80",
}

func main() {
	index := 0
	for i := 0; i < 10; i++ {
		if index >= len(devices) {
			index = 0
		}
		fmt.Println(devices[index])
		index++
	}
}
