package main

import (
	"fmt"
	"math"
)

// 最小连接
func main() {
	data := map[string]int{
		"192.168.1.2:80": 0,
		"192.168.1.3:80": 0,
		"192.168.1.4:80": 0,
	}

	var addr string
	for i := 0; i < 10; i++ {
		idx := math.MaxInt
		for k, v := range data {
			if v <= idx {
				idx = v
				addr = k
			}
		}
		fmt.Println(addr)
		data[addr]++
	}
}
