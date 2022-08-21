package mid

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sony/sonyflake"
)

// Firstly 首个执行的中间件，设置一些相关参数，做后续铺垫
func Firstly() gin.HandlerFunc {
	m := newMetrics()

	s := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Now(),
		MachineID: func() (uint16, error) {
			return 1, nil
		},
	})

	return func(c *gin.Context) {
		id, _ := s.NextID()
		c.Set(ValuesKey, &Values{
			TraceID: id,
			Now:     time.Now(),
		})
		c.Set(MetricsKey, m)
		m.Requests.Add()
		c.Next()
	}
}
