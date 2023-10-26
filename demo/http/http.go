package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	go func() {
		g := gin.New()
		g.GET("/*path", func(c *gin.Context) {
			time.Sleep(50 * time.Second)
			c.String(200, c.Request.URL.Path)
		})
		g.Run(":9999")
	}()

	r.GET("/*path", func(c *gin.Context) {

		ctx, cancel := context.WithTimeoutCause(context.Background(), 1*time.Second, fmt.Errorf("1123"))
		defer cancel()
		r := c.Request.WithContext(ctx)
		p := httputil.ReverseProxy{
			Director: func(r *http.Request) {
				r.URL.Scheme = "http"
				r.URL.Host = "192.168.10.4:9999"
				r.URL.Path = c.Request.URL.Path
			},
			Transport: &http.Transport{
				ResponseHeaderTimeout: 2 * time.Second,
			},
			ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
				w.WriteHeader(400)
				w.Write([]byte(err.Error()))
			},
		}
		h := http.TimeoutHandler(&p, 10*time.Second, "timeout")
		h.ServeHTTP(c.Writer, r)
	})
	fmt.Println("start")
	r.Run(":1123")
}
