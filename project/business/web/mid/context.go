package mid

import (
	"context"
	"time"
)

const valuesKey = "ValuesKey"

// Values 每个请求的状态
type Values struct {
	TraceID uint64
	Now     time.Time
}

// MustGetValues ...
func MustGetValues(ctx context.Context) *Values {
	m := ctx.Value(valuesKey)
	if m != nil {
		return m.(*Values)
	}
	return nil
}

// MustGetTraceID ...
func MustGetTraceID(ctx context.Context) uint64 {
	m := MustGetValues(ctx)
	if m != nil {
		return m.TraceID
	}
	return 0
}
