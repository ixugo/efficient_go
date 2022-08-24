package mid

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Logger ..
func Logger(log *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 约定: 在前置中间件中初始化了相关参数
		v := MustGetValues(c)
		c.Next()

		code := c.Writer.Status()
		if code == 200 {
			log.Infow(
				"request completed",
				"traceid", v.TraceID,
				"method", c.Request.Method,
				"path", c.Request.URL.Path,
				"remoteaddr", c.Request.RemoteAddr,
				"statuscode", code,
				"since", time.Since(v.Now).Milliseconds(),
			)
			return
		}

		// 约定: 返回给客户端的错误，记录的 key 为 responseErr
		err, _ := c.Get("responseErr")
		log.Infow(
			"request completed",
			"traceid", v.TraceID,
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"remoteaddr", c.Request.RemoteAddr,
			"statuscode", code,
			"since", time.Since(v.Now).Milliseconds(),
			"err", err,
		)
	}
}
