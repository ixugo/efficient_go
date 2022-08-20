package middleware

import (
	"fmt"
	"testing"
)

func TestMiddleware(t *testing.T) {

	c := newContext()

	c.use(func(c *context) {
		fmt.Println("1 >>>>>>.")
		// c.next()
		fmt.Println("1 end")
	})

	c.use(func(c *context) {
		fmt.Println("2 >>>>>>.")
		c.next()
		fmt.Println("2 end")
	})

	c.next()
}
