package context_test

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"
)

// 参考: https://go.dev/blog/context#TOC_3.2.
// Need a key type.
type myKey int

// Need a key value.
const key myKey = 0

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			ctx := context.WithValue(ctx, key, id)
			<-ctx.Done()
			fmt.Println("Cancelled:", id)
		}(i)
	}

	cancel()
	wg.Wait()
}

func TestDealline(t *testing.T) {
	deadline := time.Now().Add(150 * time.Millisecond)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	ch := make(chan bool, 1)

	go func() {
		// Simulate work.
		time.Sleep(200 * time.Millisecond)
		// Report the work is done.
		ch <- true
	}()

	// Wait for the work to finish. If it takes too long move on.
	select {
	case d := <-ch:
		fmt.Println("work complete", d)

	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}

func TestWithValue(t *testing.T) {
	traceID := "f47ac10b-58cc"
	ctx := context.WithValue(context.Background(), key, traceID)
	withValue(ctx)
}

func withValue(ctx context.Context) {
	if uuid, ok := ctx.Value(key).(string); ok {
		fmt.Println(uuid)
	}
}

func TestWithTimeout(t *testing.T) {
	const duration = 100 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan int, 1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(150)) * time.Millisecond)
		ch <- 1
	}()

	select {
	case d := <-ch:
		fmt.Println("work complete ", d)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}

func TestRequest(t *testing.T) {
	// 开启一个服务
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("start")
		time.Sleep(1 * time.Second)
		fmt.Println("end")
		_, _ = w.Write([]byte("OK\n"))
	})
	go func() {
		err := http.ListenAndServe(":9999", nil)
		fmt.Println(err)
	}()

	// Create a new request.
	req, err := http.NewRequest("GET", "http://localhost:9999/test", nil)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	// Create a context with a timeout of 50 milliseconds.
	// 这样的超时机制，只能停止请求方，处理请求方依然会继续完成任务。
	ctx, cancel := context.WithTimeout(req.Context(), 50*time.Millisecond)
	defer cancel()

	// Bind the new context into the request.
	req = req.WithContext(ctx)

	// Make the web call and return any error. Do will handle the
	// context level timeout.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("ERROR:", err)
		time.Sleep(3 * time.Second)
		return
	}

	// Close the response body on the return.
	defer resp.Body.Close()

	// Write the response to stdout.
	_, _ = io.Copy(os.Stdout, resp.Body)
}

func TestDeadLine(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	deadline(ctx)
}

func deadline(ctx context.Context) {
	// 截止日期时间和是否存在截止日期
	d, ok := ctx.Deadline()
	fmt.Println(ok, d)
}
