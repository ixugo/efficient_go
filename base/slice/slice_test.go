package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// Go 内置三大引用类型之一
// slice

// Go 语言可以通过 nil 与 empty 来表达式不同的意思
// 你可以返回零值切片来表示错误
// empty 切片可以表示顺利，但是没有数据
func TestNilSlice(t *testing.T) {
	// 声明未赋值的零值切片  slice == nil
	var nilSlice []string
	require.Nil(t, nilSlice)

	// 声明赋值的空切片  slice != nil
	emptySlice := []string{}
	require.NotNil(t, emptySlice)

	require.Nil(t, nilSliceF())
	require.NotNil(t, emptySliceF())
}

func nilSliceF() []string {
	return nil
}
func emptySliceF() []string {
	return make([]string, 0)
}

// TestMakeSlice 从切片获取切片
// 共用同一个数组
func TestMakeSlice(t *testing.T) {
	s1 := make([]byte, 5, 8)
	s1[0] = 'a'
	s1[1] = 'b'
	s1[2] = 'c'
	s1[3] = 'd'
	s1[4] = 'e'
	require.EqualValues(t, len(s1), 5)
	require.EqualValues(t, cap(s1), 8)

	s2 := s1[2 : 2+2]
	require.EqualValues(t, len(s2), 2)
	require.EqualValues(t, cap(s2), 8-2)
	inspectSlice("slice1", s1)
	inspectSlice("slice2", s2)
	require.EqualValues(t, &s1[2], &s2[0])
	require.EqualValues(t, &s1[3], &s2[1])

	// s2 = append(s2, 'g')
	// s2 创建的切片，使用 append 会有副作用
	// 对 s2 添加元素，导致 s1 也存在同样的函数，这是理所当然的。
	// 因为共用 同一个数组

	// 无副作用使用 append 方法 1 三下标创建切片
	s3 := s1[2:4:4]
	s3 = append(s3, 'f')
	require.EqualValues(t, len(s1), 5)
	require.EqualValues(t, len(s3), 3)
	// 无副作用使用 append 方法 2 copy
	s4 := make([]byte, len(s1))
	copy(s4, s1)

}

func inspectSlice(name string, arr []byte) {
	fmt.Printf("%s length[%d]\tcapacity[%d]\n", name, len(arr), cap(arr))
	for i := range arr {
		fmt.Printf("[%d] %p\n", i, &arr[i])
	}
}
