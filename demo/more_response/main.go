package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

// 流式传输
func main() {
	http.HandleFunc("/aaa", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Content-Type", "text/plain")

		ch := make(chan Resp, 10)
		defer close(ch)
		go func() {
			// 此处 defer... recover()...
			tick := time.NewTicker(40 * time.Millisecond)
			defer tick.Stop()
			var zeroValue Resp
			var last *Resp
			fn := func(v Resp, w io.Writer) error {
				b, _ := json.Marshal(v)
				if _, err := w.Write(b); err != nil {
					return err
				}
				w.(http.Flusher).Flush()
				return nil
			}

			for {
				select {
				case <-tick.C:
					if last != nil {
						_ = fn(*last, w)
						last = nil
					}
				case v := <-ch:
					if v != zeroValue {
						last = &v
						continue
					}
					if last != nil {
						_ = fn(*last, w)
					}
					return
				}
			}
		}()

		var resp Resp
		resp.All = 300
		ok := rand.Intn(10) == 5
		for i := 0; i <= resp.All; i++ {
			time.Sleep(10 * time.Millisecond)
			fmt.Println(i)
			if ok {
				resp.Err = fmt.Errorf("err").Error()
			}
			resp.CUR = i
			ch <- resp
			if ok {
				break
			}
		}
		time.Sleep(30 * time.Second)
	})
	http.Handle("/", http.FileServer(http.Dir("./"))) // 展示网页
	_ = http.ListenAndServe(":8888", nil)
}

type Resp struct {
	All int    `json:"all"`
	CUR int    `json:"cur"`
	Err string `json:"err"`
}
