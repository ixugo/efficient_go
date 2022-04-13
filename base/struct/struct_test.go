package struct_test

import (
	"testing"
	"unsafe"
)

type struct1 struct {
	a bool
	b int16
	c int32
}
type struct2 struct {
	a bool
	b int32
	c int32
}

type struct3 struct {
	b int64
	c float32
	a bool
}

type struct4 struct {
	a bool
	b int64
	c float32
}

// TestMemoryAlignment 结构体内存对齐
// 结构体的字段顺序，将影响内存大小
// 如果需要优化内存，请按照从大到小的顺序排列，参考 struct3
func TestMemoryAlignment(t *testing.T) {
	var s1 struct1
	t.Logf("%d\n", unsafe.Sizeof(s1)) // 8

	var s2 struct2
	t.Logf("%d\n", unsafe.Sizeof(s2)) // 12

	var s3 struct3
	t.Logf("%d\n", unsafe.Sizeof(s3)) // 16

	var s4 struct4
	t.Logf("%d\n", unsafe.Sizeof(s4)) // 24

}

type bill struct {
	flag    bool
	counter int64
}

type alice struct {
	flag    bool
	counter int64
}

func TestSameStruct(t *testing.T) {
	b := bill{counter: 11}
	a := alice{counter: 22}

	// 强制类型转换
	a = alice(b)
	t.Log(a, b)

	c := struct {
		flag    bool
		counter int64
	}{
		counter: 33,
	}
	// 无序类型转换
	a = c
	t.Log(a, c)
}
