package test

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"runtime"
	"testing"
)

func TestPost(t *testing.T) {
	pr, pw := io.Pipe()
	w := multipart.NewWriter(pw)

	go func() {
		defer pw.Close()
		defer w.Close()
		path := `/Users/xugo/Documents/efficient_go/demo/ginx/5.5del语句.mp4`
		file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		part, _ := w.CreateFormFile("file", "第三章非谓语动词.mp4")
		_, err = io.Copy(part, file)
		if err != nil {
			panic(err)
		}
		fmt.Println("send end")
	}()

	req, _ := http.NewRequest(http.MethodPost, "http://127.0.0.1:3344/video", pr)
	req.Header.Set("content-type", w.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)
	fmt.Println(">>>>>")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	fmt.Println(string(b))
	fmt.Println("end")

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("TotalAlloc = %v KB", float64(mem.TotalAlloc)/1024)
}

func TestPost2(t *testing.T) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	path := `/Users/xugo/Documents/efficient_go/demo/ginx/5.5del语句.mp4`
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	part, _ := w.CreateFormFile("file", "第三章非谓语动词.mp4")
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}
	fmt.Println("send end")
	w.Close()

	req, _ := http.NewRequest(http.MethodPost, "http://127.0.0.1:3344/video", bytes.NewReader(buf.Bytes()))
	req.Header.Set("content-type", w.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)
	fmt.Println(">>>>>")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	fmt.Println(string(b))
	fmt.Println("end")

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("TotalAlloc = %v MiB", float64(mem.TotalAlloc)/1024/1024)
}
