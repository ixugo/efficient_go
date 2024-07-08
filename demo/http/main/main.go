package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	g := gin.New()
	g.GET("/", func(c *gin.Context) {
		a := make(chan struct{}, 1)
		go func() {
			time.Sleep(3 * time.Second)
			a <- struct{}{}
		}()

		select {
		case <-c.Request.Cancel:
			fmt.Println("请求取消")
		case <-c.Request.Context().Done():
			fmt.Println("请求超时")
		case <-a:
			fmt.Println("ok")
		}
		c.String(200, "ok")
	})

	g.Run(":1234")
}
