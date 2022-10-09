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

// TestNilRange 测试对 nil 切片遍历
func TestNilRange(t *testing.T) {
	var a []int = nil
	for _, v := range a {
		fmt.Println(v)
	}
}

// TestCompareArray 长度相同，类型相同的数组之间的比较，切片不可以使用==比较
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b) // false
	// t.Log(a == c)   // 长度不等的数组不可比较
	t.Log(a == d) // true
}

// TestNewSlice 创建切片
func TestNewSlice(t *testing.T) {
	// 声明空数组
	var s0 []int
	t.Log(len(s0), cap(s0))
	// 扩容添加一个元素
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	// 声明并初始化切片
	s1 := []int{1, 2, 3, 4, 5}
	t.Log(len(s1), cap(s1))

	// make 1:类型, 2:length, 3:capacity
	// 其中 len 个元素会被初始化为 默认 0
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
}

// TestSliceAddress 发生扩容时，会创建新的数组，并将数据复制过去
func TestSliceAddress(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

// TestSliceShareMemory 切片的数据共享
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	// 数据截取后, cap 的容量根据剩余空间来计算
	t.Log(Q2, len(Q2), cap(Q2))

	// 切片并非复制了一份数据, 而是指针
	// 数据是共享的
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)

	t.Log("测试切片的长度是否会减少...") // 会, 容量和长度都发生了变化, 拼接字符串底层会创建新数组
	t.Logf("长度为 : %d, 容量为 %d 地址为 %p", len(year), cap(year), year)
	year = append(year[2:4], year[4:]...)
	t.Logf("长度为 : %d, 容量为 %d 地址为 %p", len(year), cap(year), year)
}

// 能够比较切片吗?  不能
type TestSlice struct {
	a int
	b []byte
}

func BenchmarkSlice(b *testing.B) {
	k := []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < b.N; i++ {
		k = append([]int{i}, k...)
	}
}

func BenchmarkSlice2(b *testing.B) {
	k := []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < b.N; i++ {
		t := make([]int, 1, len(k)+1)
		t[0] = i
		k = append(t, k...)
	}
}

func TestSlice2(t *testing.T) {
	s1 := make([]int, 0, 10)

	// Go 语言只有值传递
	appendFunc := func(s []int) {
		s = append(s, 10, 20, 30)
		fmt.Println(s)
	}

	fmt.Println(s1)
	appendFunc(s1)
	fmt.Println(s1)
	fmt.Println(s1[:10])
}
