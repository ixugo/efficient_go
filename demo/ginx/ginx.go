package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type CustomResponseWriter struct {
	http.ResponseWriter
}

func NewCustomResponseWriter(w http.ResponseWriter) *CustomResponseWriter {
	return &CustomResponseWriter{w}
}

func (c *CustomResponseWriter) Unwrap() http.ResponseWriter {
	return c.ResponseWriter
}

func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	cw := NewCustomResponseWriter(w)

	rc := http.NewResponseController(cw)
	if err := rc.SetReadDeadline(time.Now().Add(30 * time.Second)); err != nil {
		fmt.Println(err)
		return
	}
	// rc.SetWriteDeadline(time.Now().Add(30 * time.Second))

	file, _, err := req.FormFile("file")
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	defer file.Close()

	f, err := os.Create("filename")
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully")
}

func main() {

	g := gin.New()

	// func(ctx *gin.Context) {
	// 	fmt.Println("hello")
	// 	h, err := ctx.FormFile("file")
	// 	if err != nil {
	// 		fmt.Printf(`h, err[%s] := ctx.FormFile("file")\n`, err)
	// 	}
	// 	if err := ctx.SaveUploadedFile(h, h.Filename); err != nil {
	// 		fmt.Printf(`err[%s] := ctx.SaveUploadedFile(\n`, err)
	// 	}

	// 	ctx.String(200, "ok")
	// }

	sss := http.TimeoutHandler(http.HandlerFunc(ServeHTTP), 10*time.Second, "2")
	g.POST("/hello", gin.WrapF(ServeHTTP))

	g.POST("/upload", gin.WrapH(sss))
	svr := http.Server{
		Handler:      g,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		Addr:         ":1133",
	}

	_ = svr.ListenAndServe()

}
