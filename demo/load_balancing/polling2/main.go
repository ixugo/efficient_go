package main

import (
	"fmt"
	"math/rand"
)

var devices = []string{
	"192.168.1.2:80",
	"192.168.1.3:80",
	"192.168.1.4:80",
}

func main() {
	cached := make(map[int]string)
	for i := 0; i < 10; i++ {
		v, ok := cached[i%3]
		if ok {
			fmt.Println(v)
			continue
		}
		idx := rand.Intn(len(devices))
		v = devices[idx]
		cached[i%3] = v
		fmt.Println(v)
	}
}
