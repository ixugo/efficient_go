package main

import (
	"bytes"
	"time"
)

// 模拟把内存打满的情况
// windows 下进程不会退出
// mac 和 linux 会杀掉进程
func main() {
	buf := bytes.NewBuffer([]byte("Hello"))
	for {
		time.Sleep(time.Millisecond * 200)
		buf.Write(buf.Bytes())
	}
}
