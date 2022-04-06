package escape

import (
	"fmt"
	"testing"
)

func TestEscapePointer(t *testing.T) {
	escapePointer()
}

// escapePointer 被已经逃逸变量引用的对象，会逃逸
func escapePointer() {
	str := 5
	fmt.Println(str)
}
