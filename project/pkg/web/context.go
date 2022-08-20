package web

import (
	"context"
	"errors"
	"time"
)

type ctxKey int

const key ctxKey = 1

// Values 每个请求的状态
type Values struct {
	TraceID    string
	Now        time.Time
	StatusCode int
}

// GetValues 从上下文获取值
func GetValues(ctx context.Context) (*Values, error) {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return nil, errors.New("web value missing from context")
	}
	return v, nil
}

// GetTraceID 获取
// TODO 试试用更短的 snowid
func GetTraceID(ctx context.Context) string {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return "00000000-0000-0000-0000-000000000000"
	}
	return v.TraceID
}

// func AddSpan(ctx context.Context,spanName string) {

// }

func SetStatusCode(ctx context.Context, code int) error {
	v, err := GetValues(ctx)
	if err != nil {
		return err
	}
	v.StatusCode = code
	return nil
}
