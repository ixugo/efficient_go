package channel_test

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
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

// 测试关闭 chan 后，for 循环，总是忘记
func TestCloseChan(t *testing.T) {
	ch := make(chan int, 5)
	go func() {
		ch <- 1
		time.Sleep(50 * time.Millisecond)
		ch <- 2
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("end")
}

// TestTicker ticker.Stop() 不会关闭通道，所以不会退出 for 循环
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()
	var i int
	for v := range ticker.C {
		i++
		fmt.Println(v)
		time.Sleep(1 * time.Second)
		if i >= 2 {
			break
		}
	}
	fmt.Println("end")
}

// 当通道关闭时，ok =false
// 直接判断值是否有效，不为零值即有效
func TestChan(t *testing.T) {
	ch := make(chan int, 10)
	go func() {
		time.Sleep(1 * time.Second)
		for {
			v, ok := <-ch
			fmt.Println(ok)
			fmt.Println(v)
			// if !ok {
			// 	return
			// }
		}
	}()

	for i := 1; i < 10; i++ {
		if i == 5 {
			ch <- i
			// close(ch)
			break
		}
		ch <- i
	}
	time.Sleep(time.Second * 3)
}

func worker(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	// Acquire semaphore
	fmt.Printf("Worker %d: Waiting to acquire semaphore\n", id)
	sem <- struct{}{}

	// Do work
	fmt.Printf("Worker %d: Semaphore acquired, running\n", id)
	time.Sleep(1 * time.Second)

	// Release semaphore
	<-sem
	fmt.Printf("Worker %d: Semaphore released\n", id)
}

// TestWorker 使用 channel 限制协程数量
func TestWorker(t *testing.T) {
	nWorkers := 10                   // Total number of goroutines
	maxConcurrency := 2              // Allowed to run at the same time
	batchInterval := 3 * time.Second // Delay between each batch of 2 goros

	// Create a buffered channel with a capacity of maxConcurrency
	sem := make(chan struct{}, maxConcurrency)

	var wg sync.WaitGroup

	// We start 10 goroutines but only 2 of them will run in parallel
	for i := 1; i <= nWorkers; i++ {
		wg.Add(1)
		go worker(i, sem, &wg)

		// Introduce a delay after each batch of workers
		if i%maxConcurrency == 0 && i != nWorkers {
			fmt.Printf("Waiting for batch interval...\n")
			time.Sleep(batchInterval)
		}
	}
	wg.Wait()
	close(sem) // Remember to close the channel once done
	fmt.Println("All workers have completed")
}
