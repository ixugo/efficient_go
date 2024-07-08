package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

// 句柄泄露，会导致双方内存都暴涨。

func main() {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(30 * time.Millisecond)
			w.Write([]byte("Hello"))
		})
		http.ListenAndServe(":28081", nil)
	}()

	ch := make(chan struct{}, 10)

	for range 5000 {
		ch <- struct{}{}
		go func() {
			defer func() {
				<-ch
			}()
			resp, err := http.Get("http://localhost:28081")
			_ = err
			_ = resp
		}()
	}
	time.Sleep(50 * time.Second)
}
