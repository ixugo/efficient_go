package escape

import (
	"testing"
)

// TestBigSlice 创建超大的结合，会逃逸;
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
	// 结构体的方法内基于变量创建集合，也会逃逸
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
