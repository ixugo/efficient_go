package mid

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sony/sonyflake"
)

func SetupMetrics() gin.HandlerFunc {
	m := newMetrics()
	return func(c *gin.Context) {
		c.Set(metricsKey, m)
		m.Requests.Add()
		c.Next()
	}
}

func SetupTrace() gin.HandlerFunc {
	s := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Now(),
		MachineID: func() (uint16, error) {
			return 1, nil
		},
	})

	return func(c *gin.Context) {
		id, _ := s.NextID()
		c.Set(valuesKey, &Values{
			TraceID: id,
			Now:     time.Now(),
		})
		c.Next()
	}
}
