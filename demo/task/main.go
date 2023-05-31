package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	const numRoutines = 1000000

	var wg sync.WaitGroup
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(10 * time.Second)
			wg.Done()
		}()
	}
	wg.Wait()

	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	fmt.Printf("os:%+v\n", stats.Sys/1024/1024)
}
