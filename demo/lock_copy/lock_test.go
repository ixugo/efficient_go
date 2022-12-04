package lockcopy

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Lock struct {
	m sync.Mutex
	v int
}

// TestLock 如果是上锁状态，复制出来的也是上锁状态
func TestLock(t *testing.T) {
	var a Lock
	a.v = 5
	var wg sync.WaitGroup
	wg.Add(1)
	go func(a *Lock) {
		defer wg.Done()
		a.m.Lock()
		a.v = 3
		b := *a
		fmt.Printf("b.v:%d tryLock:%t\n", b.v, b.m.TryLock())
		time.Sleep(2 * time.Second)
		a.m.Unlock()
	}(&a)
	wg.Wait()
}
