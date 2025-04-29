package upload

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestMimeType 判断文件的 mime type
func TestMimeType(t *testing.T) {
	pr, pw := io.Pipe()
	w := multipart.NewWriter(pw)

	go func() {
		defer w.Close()

		part, err := w.CreateFormFile("file", "do.png")
		require.NoError(t, err)
		f, err := os.Open("/Users/xugo/Downloads/xjpic.do.PNG")
		require.NoError(t, err)
		defer f.Close()
		_, _ = io.Copy(part, f)
	}()

	req := httptest.NewRequest("POST", "/upload", pr)
	req.Header.Add("Content-Type", w.FormDataContentType())

	ser := server()
	rr := httptest.NewRecorder()
	ser.ServeHTTP(rr, req)
	resp := rr.Result()
	defer resp.Body.Close()
	fmt.Println("======response body=======")
	_, _ = io.Copy(os.Stdout, resp.Body)
	fmt.Println("\n==========================")
}

func server() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.WriteHeader(200)

		fmt.Println(r.Header.Get("Content-Length"))

		f, head, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(400)
			_, _ = io.WriteString(w, err.Error())
			return
		}
		defer f.Close()

		if head.Size > 1024*1024*20 {
			w.WriteHeader(400)
			fmt.Println(head.Size/1024/1024, "MB")
			_, _ = io.WriteString(w, "文件太大")
			return
		}

		var a [512]byte
		if _, err = f.Read(a[:]); err != nil {
			w.WriteHeader(400)
			return
		}

		t := http.DetectContentType(a[:])

		allowed := false
		allowedTypes := []string{"image/jpeg", "image/png", "image/gif"}
		for _, v := range allowedTypes {
			if strings.EqualFold(t, v) {
				allowed = true
				break
			}
		}

		if !allowed {
			// 不允许的类型
			w.WriteHeader(400)
			return
		}

		file, _ := os.Create(filepath.Join("/Users/xugo/Documents/efficient_go/demo/upload", head.Filename))
		defer file.Close()
		_, _ = f.Seek(0, io.SeekStart)
		_, _ = io.Copy(file, f)

		_, _ = io.WriteString(w, t)
	})
	return m
}

func TestIsNotExist(t *testing.T) {
	fmt.Println(os.IsNotExist(nil))
	fmt.Println(os.IsNotExist(errors.New("123")))
	fmt.Println(os.IsNotExist(os.ErrNotExist))

	_, err := os.Stat("/Users/xugo/Documents/efficient_go/demo/upload/upload_test.go")
	fmt.Println(err)
	fmt.Println(os.IsNotExist(err))
	fmt.Println(os.IsExist(err))

	fmt.Println(">>>>>>>>>>>>>>>>>>")

	// _, err = os.Stat("/Users/xugo/Documents/efficient_go/demo/upload/upload_test.go.go")
	// fmt.Println(err)
	// fmt.Println(os.IsNotExist(err))
	// fmt.Println(os.IsExist(err))
}
