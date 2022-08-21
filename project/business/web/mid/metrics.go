package mid

import (
	"context"
	"expvar"
)

// MetricsKey ...
const MetricsKey = "MetricsKey"

type Int interface {
	Add()
}

type expvarInt struct {
	i *expvar.Int
}

func newExpvarInt(key string) *expvarInt {
	return &expvarInt{i: expvar.NewInt(key)}
}

func (e *expvarInt) Add() {
	e.i.Add(1)
}

// Metrics 监控指标
type Metrics struct {
	Goroutines Int
	Requests   Int
	Errors     Int
	Panics     Int
}

func newMetrics() *Metrics {
	return &Metrics{
		Goroutines: newExpvarInt("goroutines"),
		Requests:   newExpvarInt("requests"),
		Errors:     newExpvarInt("errors"),
		Panics:     newExpvarInt("panics"),
	}
}

// MustGetMetrics ...
func MustGetMetrics(ctx context.Context) *Metrics {
	m := ctx.Value(MetricsKey)
	if m != nil {
		return m.(*Metrics)
	}

	return nil
}
