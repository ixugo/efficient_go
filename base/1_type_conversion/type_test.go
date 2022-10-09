package conversion

import (
	"testing"
	"unsafe"
)

// 定义别名
type MyInt int64

// Go 语言没有隐式类型转换
// 多算或少算一个字节，程序就会出问题
// 使用强制类型转换，必须清楚操作带来的影响
// 为了保证程序可靠，类型转换会增加内存开销
func TestImplicit(t *testing.T) {
	var a int32 = 1
	b := int64(a)
	c := int16(a)
	t.Logf("a: %d\t%p\t%d", a, &a, unsafe.Sizeof(a))
	t.Logf("b: %d\t%p\t%d", b, &b, unsafe.Sizeof(b))
	t.Logf("c: %d\t%p\t%d", c, &c, unsafe.Sizeof(c))
}
