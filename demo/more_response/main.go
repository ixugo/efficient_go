package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// 流式传输
func main() {
	http.HandleFunc("/aaa", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)

		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Content-Type", "text/html")

		var resp struct {
			Total   int `json:"total"`
			Current int `json:"current"`
		}

		flush := w.(http.Flusher)
		for i := 0; i <= 30; i++ {
			resp.Total = 30
			resp.Current = i
			b, _ := json.Marshal(resp)
			_, _ = w.Write(b)
			flush.Flush()
			time.Sleep(300 * time.Millisecond)
		}
	})
	http.Handle("/", http.FileServer(http.Dir("./")))
	_ = http.ListenAndServe(":8888", nil)
}
