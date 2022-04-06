package escape

import "testing"

// TestGetPointer 函数返回指针，会逃逸
func TestGetPointer(t *testing.T) {
	_ = getPointer()
}

// 逃逸分析，函数返回指针类型，会逃逸
func getPointer() *int {
	a := 5
	return &a
}
