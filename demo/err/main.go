package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		var out struct {
			Reason  string   `json:"reason"`  // 给机器看的(某些场景不一定需要)
			Msg     string   `json:"msg"`     // 给用户看的
			Details []string `json:"details"` // 给技术工作者看的
		}
		out.Reason = "Unknow"
		out.Msg = "读取 /home/app/main.jsx 文件失败"
		out.Details = append(out.Details, "since=12451243123", "尝试先创建文件", `Download the React DevTools for a better development experience: https://reactjs.org/link/react-devtools`)
		b, _ := json.Marshal(out)
		_, _ = w.Write(b)
	})

	http.Handle("/", http.FileServer(http.Dir(".")))
	_ = http.ListenAndServe(":8888", nil)
}
