package main

import (
	"fmt"
	"io"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	_ = trace.Start(os.Stdout)
	defer trace.Stop()

	for i := 0; i < 100000; i++ {
		fmt.Fprintf(io.Discard, fmt.Sprintf("%s %d", "123", 1))
		time.Sleep(1 * time.Microsecond)
	}
}
