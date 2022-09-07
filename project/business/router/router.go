package router

import (
	"os"
	"project/business/web/mid"
	"project/pkg/convert"

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
		mid.Limiter(200),
	)

	gin.Default()

	r.GET("/test", test)

	users := r.Group("/users", mid.TokenHandle("key"))
	users.GET("/:id", func(ctx *gin.Context) {
		id := convert.StrTo(ctx.Param("id")).MustInt()
		ctx.JSON(200, gin.H{
			"token": mid.NewToken(id, "key"),
		})
	})
	return r
}

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"program": os.Args[0],
	})
}
