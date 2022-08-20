package service

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	return e
}
