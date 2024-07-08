package network

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"syscall"
	"testing"
	"time"
)

func TestListener(t *testing.T) {
	// 0 表示随机端口
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	defer lis.Close()

	go func() {
		for {
			conn, err := lis.Accept()
			if err != nil {
				panic(err)
			}
			go func(conn net.Conn) {
				defer conn.Close()
				fmt.Println(conn.RemoteAddr())
			}(conn)
		}
	}()
	time.Sleep(10 * time.Second)
	t.Logf("bound to %q", lis.Addr())
}

func TestDial(t *testing.T) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	fmt.Println("启动本地监听", lis.Addr())
	done := make(chan struct{})
	go func() {
		defer func() { done <- struct{}{} }()

		for {
			conn, err := lis.Accept()
			if err != nil {
				t.Log("连接不上", err)
				break
			}
			conn.(*net.TCPConn).SetReadBuffer(1024 * 1024)

			go func(c net.Conn) {
				defer func() {
					c.Close()
					done <- struct{}{}
				}()
				fmt.Println("来了连接", c.RemoteAddr())
				c.SetReadDeadline(time.Now().Add(50 * time.Millisecond))

				buf := make([]byte, 1024)
				for {
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.Error("发生错误", err)
						}
						// 大部分需要处理的错误是超时错误
						// https://pkg.go.dev/net#OpError/
						nErr, ok := err.(net.Error)
						if ok && nErr.Timeout() {
							fmt.Println("读取超时了")
						}

						if err, ok := err.(*net.OpError); ok {
							if err.Timeout() {
								fmt.Println("读取超时了")
							}
							if err.Temporary() {
								fmt.Println("资源超限")
							}
						}
						fmt.Println("连接断开", c.RemoteAddr())
						return
					}
					t.Logf("received: %q", buf[:n])
				}
			}(conn)
		}
	}()

	conn, err := net.Dial("tcp", lis.Addr().String())
	if err != nil {
		t.Fatal("连接不上", err)
	}
	time.Sleep(time.Second)
	conn.Write([]byte("hello"))

	conn.Close()
	<-done
	time.Sleep(2 * time.Second)
	lis.Close()
	<-done
}

func TestDialTimeout(t *testing.T) {
	now := time.Now()

	conn, err := net.DialTimeout("tcp", "10.0.0.0:9902", 5*time.Second)
	if err != nil {
		fmt.Println(time.Since(now))
		fmt.Println(err)
		return
	}
	_ = conn
}

func TestDialWithContext(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	var d net.Dialer
	d.Control = func(network, address string, c syscall.RawConn) error {
		time.Sleep(5 * time.Second * time.Millisecond)
		return nil
	}
	conn, err := d.DialContext(ctx, "tcp", "10.0.0.0:9902")
	if err != nil {
		panic(err)
	}
	_ = conn
}

// 并发 dial
// 以下模拟 dial 并发连接，服务器仅接口一个时，其它并发的停止连接。
func TestMultiple(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	defer lis.Close()

	go func() {
		// 仅运行一个连接
		conn, err := lis.Accept()
		if err == nil {
			conn.Close()
		}
	}()

	dial := func(ctx context.Context, address string, response chan int,
		id int, wg *sync.WaitGroup,
	) {
		defer wg.Done()

		var d net.Dialer
		c, err := d.DialContext(ctx, "tcp", address)
		if err != nil {
			return
		}
		c.Close()

		select {
		case <-ctx.Done():
		case response <- id:
		}
	}

	res := make(chan int)
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go dial(ctx, lis.Addr().String(), res, i+1, &wg)
	}

	response := <-res
	cancel()
	wg.Wait()
	close(res)

	if ctx.Err() != context.Canceled {
		t.Fatal(ctx.Err())
	}

	t.Logf("dialer %d", response)
}

func TestDeadline(t *testing.T) {
	sync := make(chan struct{})
	lis, err := net.Listen("tcp", ":")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		conn, err := lis.Accept()
		if err != nil {
			t.Log(err)
			return
		}
		defer func() {
			conn.Close()
			close(sync)
		}()
		if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
			t.Error(err)
			return
		}
		buf := make([]byte, 1)
		_, err = conn.Read(buf) // blocked until remote node sends data
		nErr, ok := err.(net.Error)
		if !ok || !nErr.Timeout() {
			t.Error("exported timeout")
		}

		sync <- struct{}{}

		if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
			t.Error(err)
			return
		}

		_, err = conn.Read(buf)
		if err != nil {
			t.Error(err)
		}
	}()

	conn, err := net.Dial("tcp", lis.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	<-sync

	if _, err := conn.Write([]byte("1")); err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 1)
	if _, err := conn.Read(buf); err != io.EOF {
		t.Error("exported EOF,buf actual", err)
	}
}

const defaultPingInterval = 30 * time.Second

func TestHeartbeat(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	r, w := io.Pipe()
	done := make(chan struct{})
	resetTimer := make(chan time.Duration, 1)
	resetTimer <- time.Second

	go func() {
		Pinger(ctx, w, resetTimer)
		close(done)
	}()

	receivePing := func(d time.Duration, r io.Reader) {
		if d >= 0 {
			fmt.Printf("reset timer (%s)\n", d.String())
			resetTimer <- d
		}

		now := time.Now()
		buf := make([]byte, 1024)
		n, err := r.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("received %q (%s)\n", buf[:n], time.Since(now).Round(100*time.Millisecond))
	}

	for i, v := range []int64{0, 200, 300, 0, -1, -1, -1} {
		fmt.Printf("run %d:\n", i+1)
		receivePing(time.Duration(v)*time.Millisecond, r)
	}
	cancel()
	<-done
}

func Pinger(ctx context.Context, w io.Writer, reset <-chan time.Duration) {
	var internal time.Duration
	select {
	case <-ctx.Done():
		return
	case internal = <-reset:
	default:
	}
	if internal <= 0 {
		internal = defaultPingInterval
	}

	timer := time.NewTimer(internal)
	defer func() {
		if !timer.Stop() {
			<-timer.C
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case newInterval := <-reset:
			// 停止计时器，返回 false 表示计时器已经停止过了，此时需要将 channnel 清空
			if !timer.Stop() {
				<-timer.C
			}
			if newInterval > 0 {
				internal = newInterval
			}
		case <-timer.C:
			if _, err := w.Write([]byte("ping")); err != nil {
				return
			}
		}
		_ = timer.Reset(internal)
	}
}

// 利用心跳提前期限
func TestPingerAdvanceDealine(t *testing.T) {
	done := make(chan struct{})
	lis, err := net.Listen("tcp", ":")
	if err != nil {
		t.Fatal(err)
	}

	begin := time.Now()
	go func() {
		defer func() {
			fmt.Println(">>>>>>. end")
			close(done)
		}()
		conn, err := lis.Accept()
		if err != nil {
			t.Log(err)
			return
		}
		// ctx, cancel := context.WithCancel(context.Background())
		defer func() {
			// cancel()
			conn.Close()
		}()
		resetTimer := make(chan time.Duration, 1)
		resetTimer <- time.Second
		// 启动一个接受连接的侦听器，将 Pinger 设置为每秒 ping 一次，并将初始截止时间设置为五秒
		// go Pinger(ctx, conn, resetTimer)
		if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
			t.Fatal(err)
			return
		}
		// conn.SetWriteDeadline()
		buf := make([]byte, 1024)
		for {
			if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
				t.Error(err)
				return
			}
			n, err := conn.Read(buf)
			if err != nil {
				return
			}
			t.Logf("服务端收到消息: [%s] %s", time.Since(begin).Truncate(time.Second), buf[:n])
			// resetTimer <- 0
		}
	}()
	conn, err := net.Dial("tcp", lis.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// buf := make([]byte, 1024)
	for range 10 {
		if _, err := conn.Write([]byte("PONG!!!")); err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second)

		// n, err := conn.Read(buf)
		// if err != nil {
		// 	t.Fatal(err)
		// }
		// t.Logf(">>> [%s] %s", time.Since(begin).Truncate(time.Second), buf[:n])
	}
	// if _, err := conn.Write([]byte("PONG!!!")); err != nil {
	// 	t.Fatal(err)
	// }

	// for range 4 {
	// 	n, err := conn.Read(buf)
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			t.Fatal(err)
	// 		}
	// 		break
	// 	}
	// 	t.Logf("[%s] %s", time.Since(begin).Truncate(time.Second), buf[:n])
	// }
	<-done
	end := time.Since(begin).Truncate(time.Second)
	t.Logf("[%s] done", end)
	if end != 9*time.Second {
		t.Fatalf("expected EOF at 9 seconds; actual %s", end)
	}
}
