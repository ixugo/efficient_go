package mid

import (
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// Panics recovers from panics and converts the panic to an error so it is
// reported in Metrics and handled in Errors.
func Panics() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Defer a function to recover from a panic and set the err return
		// variable after the fact.
		defer func() {
			if rec := recover(); rec != nil {

				// Stack trace will be provided.
				trace := debug.Stack()
				err := fmt.Errorf("PANIC [%v] TRACE[%s]", rec, string(trace))
				fmt.Println(err)

				v := MustGetMetrics(c)
				v.Panics.Add()

				// Updates the metrics stored in the context.
				// metrics.AddPanics(ctx)
			}
		}()

		c.Next()
	}

}
