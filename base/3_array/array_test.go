package array

import (
	"fmt"
	"testing"
)

// TestArrayInit 定义数组
func TestArrayInit(t *testing.T) {
	// 定义数组后自动全部初始化为 0
	var arr [3]int
	t.Log(arr[1], arr[2])

	// 定义数组长度并赋值
	arr1 := [4]int{1, 2, 3, 4}

	// 默认数组长度并赋值,值的数量等于最终长度
	arr2 := [...]int{1, 2, 3, 4, 5}
	t.Log(arr1[1], arr2[2])
}

// 遍历数组
func TestRangeArray(t *testing.T) {
	arr2 := [...]int{1, 2, 3, 4, 5}

	// 获取下标和值
	for idx, ele := range arr2 {
		t.Log(idx, ele)
	}

	// 仅关注值
	for _, ele := range arr2 {
		t.Log(ele)
	}

	// 仅关注下标
	for i := range arr2 {
		t.Log(i)
	}
}

// 切片数组
func TestSlice(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}

	// 前 3
	arrSec := arr[:3]

	t.Log(arrSec)
	t.Log(len(arr))
	// 从 下标 2 截取到 下标 5-1
	t.Log(arr[2:])

	s := "哈哈哈"
	fmt.Println(len([]rune(s)))
}

// TestArrayOfType 多类型数组, 接口零值为 nil
func TestArrayOfType(t *testing.T) {
	var arr [5]interface{}
	arr[0] = 1
	arr[1] = "123"
	arr[2] = 3.13

	for _, value := range arr {
		fmt.Println(value)
	}
}

func TestArray(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	b := make([]*int, len(s))
	// i =>  index
	// v =>  value
	for i, v := range s {
		b[i] = &v
	}
	fmt.Println(b)
}
