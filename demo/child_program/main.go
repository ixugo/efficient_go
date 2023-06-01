package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

// 进程间通信

func main() {
	cmd := exec.Command("/Users/xugo/Documents/efficient_go/demo/child_program/child/main")
	w, _ := cmd.StdinPipe()
	r, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		fmt.Println("启动子进程失败：", err)
		return
	}
	fmt.Println("start")
	go func() {
		_, _ = w.Write([]byte("Hello\n"))
		time.Sleep(time.Second)
		_, _ = w.Write([]byte("world\n"))
		time.Sleep(time.Second)
		_, _ = w.Write([]byte("Go 是最好的语言\n"))
		_, _ = w.Write([]byte("没有之一\n"))
		time.Sleep(3 * time.Second)
		_ = cmd.Process.Kill()
	}()
	read := bufio.NewReader(r)
	for {
		l, _, err := read.ReadLine()
		if err != nil {
			return
		}
		// 忽略无意义的内容
		var data map[string]any
		if err := json.Unmarshal(l, &data); err != nil {
			fmt.Println("无意义内容", string(l))
			// 无意义的内容
			continue
		}

		fmt.Println("子进程输出：", data)
	}
}
