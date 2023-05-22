package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	// r.Use(Limiter())
	r.GET("/test", func(ctx *gin.Context) {
		time.Sleep(2 * time.Second)
		ctx.JSON(200, gin.H{"msg": "SUCCESS"})
	})

	filePath := filepath.Join(filepath.Dir(os.Args[0]), "video")

	r.GET("/video/:id", func(c *gin.Context) {
		id := c.Param("id")
		video, err := os.Open(filepath.Join(filePath, id+".mp4"))
		if err != nil {
			c.String(400, err.Error())
			return
		}
		defer video.Close()
		c.Writer.Header().Add("content-type", "video/mp4")
		// 作为二进制流返回
		http.ServeContent(c.Writer, c.Request, "", time.Now(), video)

		// 升级云服务，伪代码
		// 鉴权，并重定向到 oss
		// 可以对 oss 访问加上有时间限制的 token
		// http.Redirect(c.Writer, c.Request, "oss", 301)
	})

	r.POST("/video", func(c *gin.Context) {
		// 限制上传最大 50 m
		const size = 1024 * 1024 * 50
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, size)
		// ParseMultipartForm 限制文件大小
		if err := c.Request.ParseMultipartForm(size); err != nil {
			c.String(400, err.Error())
			return
		}

		fileHeader, err := c.FormFile("file")
		if err != nil {
			c.String(400, err.Error())
			return
		}
		file, err := fileHeader.Open()
		if err != nil {
			c.String(400, err.Error())
			return
		}
		defer file.Close()

		f, err := os.Create(filepath.Join(filePath, fileHeader.Filename))
		if err != nil {
			c.String(400, err.Error())
			return
		}
		_, _ = io.Copy(f, file)
		defer f.Close()

		c.String(http.StatusCreated, "ok")

		// 升级云服务，伪代码
		// 上传到阿里云
	})

	r.Run(":3344")
}
