package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	firstPid := os.Getpid()
	fmt.Println(firstPid)

	var i int

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("fork")
		pid, err := os.StartProcess("./fork", nil, &os.ProcAttr{})
		if err != nil {
			panic(err)
		}
		fmt.Println(pid)
		os.Exit(0)
		// p, err := os.FindProcess(firstPid)
		// if err != nil {
		// 	panic(err)
		// }
		// p.Kill()
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		i++
		w.WriteHeader(200)
		io.WriteString(w, fmt.Sprintf("Hello%d", i))
	})
	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Println(err)
	}
}
