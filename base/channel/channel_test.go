package channel_test

import (
	"bytes"
	"runtime"
	"strconv"
	"testing"
	"time"
)

// Go 内置三大引用类型之一
// channel

// TestCloseChannel 向已关闭的 channel 发送数据，会 panic
func TestCloseChannel(t *testing.T) {
	const count = 3
	ch := make(chan string, count)
	close(ch)
	go func() {
		ch <- "end"
	}()

	time.Sleep(1 * time.Second)
}

// TestOneDone 仅需任意任务完成
func TestOneDone(t *testing.T) {

	t.Log(runtime.NumGoroutine())

	const count = 3
	// 注意使用有缓冲 channel，防止泄露
	// 尝试将 第二个参数改成 0，查看 G 的数量
	ch := make(chan string, count)
	for i := 0; i < count; i++ {
		go func(i int) {
			time.Sleep(1 * time.Second)
			ch <- "param: " + strconv.Itoa(i)
		}(i)
	}

	t.Log(<-ch)
	t.Log(runtime.NumGoroutine())
}

// TestAllDone 完成所有任务
func TestAllDone(t *testing.T) {
	t.Log(runtime.NumGoroutine())

	const count = 3
	ch := make(chan string, count)
	for i := 0; i < count; i++ {
		go func(i int) {
			time.Sleep(100 * time.Millisecond)
			ch <- "param: " + strconv.Itoa(i)
		}(i)
	}

	result := bytes.NewBufferString("\n")

	for i := 0; i < count; i++ {
		result.WriteString(<-ch + "\n")
	}

	t.Log(result.String())
	t.Log(runtime.NumGoroutine())
}
