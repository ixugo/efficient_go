package base

import (
	"testing"
)

func TestMath1(t *testing.T) {
	// a<=b，则 a&b == a
	a := 1
	b := 15

	if a > b {
		a = b
	}
	t.Log(a & b)

	// 其它写法
	c := a
	if a > b {
		c = b
	}
	t.Log(c)
}
