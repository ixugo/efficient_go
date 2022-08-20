package middleware

import "fmt"

type handler func(c *context)

type context struct {
	handlers []handler
	index    uint8
}

func newContext() *context {
	h := make([]handler, 0, 10)
	h = append(h, func(c *context) {
		c.next()
	})
	return &context{
		handlers: h,
	}
}

func (c *context) use(h handler) {
	c.handlers = append(c.handlers, h)
}

func (c *context) next() {
	fmt.Println(">>>>>. next : ", c.index)
	c.index++
	for c.index < uint8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
	fmt.Println(">>>>>. end : ", c.index)
}
