package router

import (
	"os"
	"project/business/web/mid"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRouter(log *zap.SugaredLogger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(
		mid.Firstly(),
		mid.Logger(log),
		mid.Panics(),
	)

	gin.Default()

	r.GET("/test", test)
	return r
}

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"program": os.Args[0],
	})
}
