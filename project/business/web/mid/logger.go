package mid

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Logger ..
func Logger(log *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		v := MustGetValues(c)
		c.Next()
		log.Infow(
			"request completed",
			"traceid", v.TraceID,
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"remoteaddr", c.Request.RemoteAddr,
			"statuscode", v.StatusCode,
			"since", time.Since(v.Now).Milliseconds(),
		)

	}
}
