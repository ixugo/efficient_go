package escape

import (
	"testing"
)

// go test -gcflags "-m" ./slice_test.go

// TestBigSlice 创建超大的组合，会逃逸;
func TestBigSlice(t *testing.T) {
	s := make([]int, 0, 1024*9)
	for index, _ := range s {
		s[index] = index
	}
}

// TestSlice 基于变量创建的集合，会逃逸
func TestSliceCrete(t *testing.T) {
	c := 5
	_ = make([]int, c)
	u := User{5}
	u.Handle()
}

type User struct {
	Size int
}

func (u *User) Handle() {
	arr := make([]int, u.Size)
	arr[0] = 5
}
