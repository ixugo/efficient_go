package mid

import (
	"context"
	"time"
)

const ValuesKey = "ValuesKey"

// Values 每个请求的状态
type Values struct {
	TraceID    uint64
	Now        time.Time
	StatusCode int
}

func MustGetValues(ctx context.Context) *Values {
	m := ctx.Value(ValuesKey)
	if m != nil {
		return m.(*Values)
	}
	return nil
}
