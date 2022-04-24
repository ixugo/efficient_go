package logger

import (
	"fmt"
	"io"
	"sync"
)

// Logger ...
type Logger struct {
	ch chan string
	wg sync.WaitGroup
}

// New ...
func New(w io.Writer, cap int) *Logger {
	l := Logger{
		ch: make(chan string, cap),
	}

	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		for v := range l.ch {
			fmt.Fprintln(w, v)
		}
	}()
	return &l
}

// Close ...
func (l *Logger) Close() {
	close(l.ch)
	l.wg.Wait()
}

// Println ....
func (l *Logger) Write(v string) {
	select {
	case l.ch <- v:
	default:
		fmt.Println("Drop ", v)
	}
}
