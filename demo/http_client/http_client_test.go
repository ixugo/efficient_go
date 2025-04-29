package httpclient

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"syscall"
	"testing"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func TestLimit(t *testing.T) {
	var rlimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit)
	if err != nil {
		fmt.Println("Error getting RLIMIT_NOFILE:", err)
		return
	}

	fmt.Printf("Soft limit: %d\n", rlimit.Cur)
	fmt.Printf("Hard limit: %d\n", rlimit.Max)
}

// 通过限制 MaxConnsPerHost 来限制并发数量

func TestHttpClient(t *testing.T) {
	cli := http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    10,
			MaxConnsPerHost: 1000, // 最大并发主机连接数
		},
	}

	go func() {
		var wg sync.WaitGroup
		for range 1000000000 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				resp, err := cli.Get("http://127.0.0.1:8534/test")
				if err != nil {
					fmt.Println(err)
					return
				}
				defer resp.Body.Close()
			}()
		}
		wg.Wait()
	}()

	var i atomic.Int32
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		old := i.Add(1)
		fmt.Println(old)
		time.Sleep(5 * time.Second)
		w.WriteHeader(200)
	})
	http.ListenAndServe(":8534", nil)
}

func TestH2C(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "h2c")
		time.Sleep(10 * time.Second)
	})
	s := http.Server{
		Addr:    ":33536",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}
	go func() {
		fmt.Println(s.ListenAndServe())
	}()
	go fmt.Println("123")

	fmt.Println(">>>>>>")
	client := http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.DialTimeout(network, addr, 2*time.Second)
			},
		},
	}
	var wg sync.WaitGroup

	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := client.Get("http://127.0.0.1:33536/test")
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(resp.Proto)
		}()
	}
	wg.Wait()
	fmt.Println("end")
}
