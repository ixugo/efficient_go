package func_test

import (
	"testing"
)

// TestStackCopy
func TestStackCopy(t *testing.T) {
	s := "HELLOW"
	stackCopy(&s, 0, [size]int{})
}

const size = 1000

// stackCopy 栈扩容及拷贝 (扩容 25%)
func stackCopy(s *string, count int, arr [size]int) {
	println(count, s, *s)
	count++
	if count == 10 {
		return
	}
	stackCopy(s, count, arr)
}
