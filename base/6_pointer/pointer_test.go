package pointer

import (
	"fmt"
	"testing"
)

func TestPirnt(t *testing.T) {
	a := 1
	b := &a
	fmt.Printf("%d %p\n", a, &b)
	print(b)
}

func print(v *int) {
	fmt.Printf("%d %p", *v, &v)
}

func TestUpdateSlice(t *testing.T) {
	arr := make([]int, 0, 2)
	arr = append(arr, 1)
	arr = append(arr, 2)
	fmt.Println("函数执行前: ", arr)
	updateSlice(arr)
	fmt.Println("函数执行后: ", arr)
}

func updateSlice(v []int) {
	v[0] = 9
}

func TestAppendSlice(t *testing.T) {
	arr := make([]int, 0, 2)
	arr = append(arr, 1)
	arr = append(arr, 2)
	fmt.Println("函数执行前: ", arr)
	appendSlice(arr)
	fmt.Println("函数执行后: ", arr)
}

func appendSlice(v []int) {
	v = append(v, 1, 2, 3, 4, 5)
	fmt.Println(v)
}
