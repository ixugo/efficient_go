package isync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 互斥锁 sync.Mutex 通常用来保护临界区和共享资源
// 条件变量 sync.Cond 用来协调想要访问共享资源的 goroutine。
type Cond struct {
	*sync.Cond
	done bool
}

func TestCond(t *testing.T) {
	cond := Cond{
		Cond: sync.NewCond(new(sync.Mutex)),
	}
	go cond.read("r1")
	go cond.read("r2")
	go cond.read("r3")

	cond.write("write")
	time.Sleep(time.Second)
}

func (c *Cond) read(name string) {
	c.L.Lock()
	fmt.Println(name, "lock")
	for !c.done {
		fmt.Println(name, "wait")
		c.Wait()
	}
	fmt.Println(name, "starts reading")
	c.L.Unlock()
}

func (c *Cond) write(name string) {
	fmt.Println(name, "starts writing")
	time.Sleep(time.Second)
	c.L.Lock()
	c.done = true
	c.L.Unlock()
	fmt.Println(name, "wakes all")
	// 唤醒所有
	// c.Broadcast()
	c.Broadcast()
}
