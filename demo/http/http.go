package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/*path", func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path)
		p := httputil.ReverseProxy{
			Director: func(r *http.Request) {
				r.URL.Scheme = "http"
				r.URL.Host = "127.0.0.1:9999"
				r.URL.Path = c.Request.URL.Path
			},
		}

		p.ServeHTTP(c.Writer, c.Request)
	})
	fmt.Println("start")
	r.Run(":1123")

}
