package test

import (
	"fmt"
	"io"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestConcurrent(t *testing.T) {
	type User struct {
		Name string
		Age  int32
	}
	u := &User{Name: "1", Age: 11}
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				time.Sleep(50 * time.Millisecond)
				fmt.Fprintf(io.Discard, "%d\t%s\n", i, u.Name)
			}
		}(i)
	}
	i := 0
	for {
		i++
		if i > 10 {
			break
		}
		time.Sleep(100 * time.Millisecond)
		*u = User{Name: strconv.Itoa(i), Age: u.Age}
	}
}

func TestConcurrent1(t *testing.T) {
	type User struct {
		Name string
		Age  int32
	}
	u := &User{Name: "1", Age: 11}

	var m sync.RWMutex
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				time.Sleep(50 * time.Millisecond)
				m.RLock()
				fmt.Fprintf(io.Discard, "%d\t%s\n", i, u.Name)
				m.RUnlock()
			}
		}(i)
	}
	i := 0
	for {
		i++
		if i > 10 {
			break
		}
		time.Sleep(100 * time.Millisecond)
		m.Lock()
		*u = User{Name: strconv.Itoa(i), Age: u.Age}
		m.Unlock()
	}
}

func TestConcurrent2(t *testing.T) {
	type User struct {
		Name string
		Age  int32
	}
	var u atomic.Value
	u.Store(&User{Name: "1", Age: 11})
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				time.Sleep(50 * time.Millisecond)
				user := u.Load().(*User)
				fmt.Fprintf(io.Discard, "%d\t%s\n", i, user.Name)
			}
		}(i)
	}
	i := 0
	for {
		i++
		if i > 10 {
			break
		}
		time.Sleep(100 * time.Millisecond)
		u.Store(&User{Name: strconv.Itoa(i)})
	}
}

func TestConcurrent3(t *testing.T) {
	type User struct {
		Name string
		Age  int32
	}
	var u atomic.Pointer[User]
	u.Store(&User{Name: "1", Age: 11})
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				time.Sleep(50 * time.Millisecond)
				user := u.Load()
				fmt.Fprintf(io.Discard, "%d\t%s\n", i, user.Name)
			}
		}(i)
	}
	i := 0
	for {
		i++
		if i > 10 {
			break
		}
		time.Sleep(100 * time.Millisecond)
		u.Store(&User{Name: strconv.Itoa(i)})
	}
}

func RetryWrapper(f func() error) {
	const maxDelay = 5 * time.Minute
	delay := 5 * time.Second // Initial delay
	for {
		if err := f(); err == nil {
			return
		}
		time.Sleep(delay)
		delay = min(2*delay, maxDelay)
	}
}

func min(a, b time.Duration) time.Duration {
	if a > b {
		return b
	}
	return a
}

func TestMain1(t *testing.T) {
	main1()

}

func main1() {
	var wg sync.WaitGroup
	var counter int

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			temp := counter
			temp++
			counter = temp
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			temp := counter
			temp++
			counter = temp
		}
	}()
	wg.Wait()
	fmt.Println("Counter:", counter)
}
