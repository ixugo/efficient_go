package trace

import (
	"fmt"
	"io"
)

// Tracer 描述代码中的流程事件
type Tracer interface {
	Trace(...any)
}

// New creates a new Tracer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

var _ Tracer = new(tracer)
var _ Tracer = new(nilTracer)

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...any) {
	fmt.Fprintln(t.out, a...)
}

type nilTracer struct {
}

// Trace for a nil tracer does nothing.
func (*nilTracer) Trace(...any) {
}

// Off creates a Tracer that will ignore calls to Trace.
func Off() Tracer {
	return new(nilTracer)
}
