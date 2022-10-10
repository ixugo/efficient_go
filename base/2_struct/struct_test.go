package struct_test

import (
	"fmt"
	"reflect"
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
// 如果不做内存对齐会怎样? 如果某个值的长度为 2 个字节，刚好占据在内存边界的左右两侧，于是对硬件必须做两次操作，才能读出这个数据
// 写入也是同样，这样操作的效率不高
//
// 另外需要注意的是，在考虑内存消耗之前，要确定是否有这样的需求?
// 你的程序性能瓶颈是否在这里结构体上?
// 如果不是，更建议以原有的顺序排列，更容易读懂，将意思彼此相关的属性放在一起
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

type balance struct {
	state   bool
	counter int64
}

type account struct {
	state   bool
	counter int64
}

// 两个属性相同的结构体互相赋值
func TestSameStruct(t *testing.T) {
	// var s1 struct1
	// var s2 struct2
	// if s1==s2 {}

	b := balance{counter: 11}
	a := account{counter: 22}

	// 强制类型转换
	a = account(b)
	t.Logf("\na:%v,%s\nb:%v,%s", a, reflect.TypeOf(a), b, reflect.TypeOf(b))

	// 匿名结构体，仅使用一次，没必要命名
	c := struct {
		state   bool
		counter int64
	}{
		counter: 33,
	}
	// 无需类型转换
	a = c
	t.Logf("\na:%v,%s\nb:%v,%s", a, reflect.TypeOf(a), c, reflect.TypeOf(c))
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func TestShowB(t *testing.T) {
	teacher := Teacher{}
	teacher.ShowB()
	teacher.ShowA()
}
