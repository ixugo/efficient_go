package main

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	a := make(chan int, 10)
	a <- 1
	fmt.Println(len(a), cap(a))
}
