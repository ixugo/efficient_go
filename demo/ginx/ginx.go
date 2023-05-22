package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

type ReaderL interface {
	SetReadDeadline(deadline time.Time) error
}

func ServeHTTP1(w http.ResponseWriter, req *http.Request) {

	rc := http.NewResponseController(w)
	_ = rc.SetReadDeadline(time.Now().Add(30 * time.Second))
	_ = rc.SetWriteDeadline(time.Now().Add(30 * time.Second))
	// rc.SetWriteDeadline(time.Now().Add(30 * time.Second))

	file, head, err := req.FormFile("file")
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	defer file.Close()

	f, err := os.Create(filepath.Join("./", head.Filename))
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

func ServeHTTP2(w http.ResponseWriter, req *http.Request) {

	// rc.SetWriteDeadline(time.Now().Add(30 * time.Second))

	file, head, err := req.FormFile("file")
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	defer file.Close()

	f, err := os.Create(filepath.Join("./", head.Filename))
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
	// h := g.Handler()

	// m := http.NewServeMux()

	// m.ServeHTTP(w, r)

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

	g.POST("/hello", gin.WrapF(ServeHTTP1))
	g.POST("/upload", gin.WrapF(ServeHTTP2))

	svr := http.Server{
		Handler:      g,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		Addr:         ":1133",
	}
	_ = svr.ListenAndServe()

}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := http.NewServeMux()
	m.ServeHTTP(w, r)

}
