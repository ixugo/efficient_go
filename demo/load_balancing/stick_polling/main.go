package main

import (
	"fmt"
	"math/rand"
)

func main() {
	data := map[string]int{
		"192.168.1.2:80": 4,
		"192.168.1.3:80": 3,
		"192.168.1.4:80": 3,
	}

	ds := make([]string, 0, 10)
	for k, v := range data {
		for i := 0; i < v; i++ {
			ds = append(ds, k)
		}
	}

	for i := 0; i < 10; i++ {
		v := rand.Intn(len(ds))
		fmt.Println(ds[v])
	}
}
