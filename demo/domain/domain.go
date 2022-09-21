package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/live", func(ctx *gin.Context) {
		fmt.Println(ctx.Request.Host)
		ctx.JSON(200, gin.H{"msg": "OK"})
	})

	r.Any("/test/*path", func(ctx *gin.Context) {
		ctx.String(200, ctx.Request.URL.Path)
	})
	r.Run(":9999")
}
