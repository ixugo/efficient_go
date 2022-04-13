package slice

import (
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
