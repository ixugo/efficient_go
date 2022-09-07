package mid

import (
	"context"
	"expvar"
)

// metricsKey ...
const metricsKey = "MetricsKey"

type inter interface {
	Add()
}

type expvarInt struct {
	i *expvar.Int
}

var e = make(map[string]struct{}, 4)

func newExpvarInt(key string) *expvarInt {
	_, exist := e[key]
	if exist {
		panic("metrics 存在相同的 Key")
	}
	return &expvarInt{i: expvar.NewInt(key)}
}

func (e *expvarInt) Add() {
	e.i.Add(1)
}

// Metrics 监控指标
type Metrics struct {
	Goroutines inter
	Requests   inter
	Errors     inter
	Panics     inter
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
	m := ctx.Value(metricsKey)
	if m != nil {
		return m.(*Metrics)
	}

	return nil
}
