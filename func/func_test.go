package func_test

import (
	"fmt"
	"testing"
)

func TestEscapes(t *testing.T) {
	_ = getInt()
	myPrint()
	maps()

}

// TestSlice 基于变量创建的集合，会逃逸
func TestSlice(t *testing.T) {
	c := 5
	_ = make([]int, c)
	// 结构体的方法内基于变量创建集合，也会逃逸
	u := User{5}
	u.Handle()
}

// 逃逸分析，函数返回指针类型，会逃逸
func getInt() *int {
	a := 5
	return &a
}

func maps() {
	tmp := make(map[string]int, 1000)
	fmt.Printf("%T", tmp)
	_ = tmp
}

type User struct {
	Size int
}

func (u *User) Handle() {
	arr := make([]int, u.Size)
	arr[0] = 5
}

// 被已经逃逸变量引用的对象，会逃逸
func myPrint() {
	str := 5
	fmt.Println(str)
}

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
