package escape

import (
	"fmt"
	"testing"
)

// TestGetPointer 函数返回指针，会逃逸
// go test -gcflags "-m" ./func_test.go
func TestGetPointer(t *testing.T) {
	_ = getPointer()
}

// 逃逸分析，函数返回指针类型，会逃逸
func getPointer() *int {
	a := 5
	return &a
}

func TestGetPointer1(t *testing.T) {
	_ = getPointer()
}

func TestDefer(t *testing.T) {
	deferFunc()

}

func deferFunc() {
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		a := 0
		if a == 0 {
			defer func() {
				fmt.Println("2")
			}()
		}
		fmt.Println("3")
	}()
	fmt.Println("4")
}
