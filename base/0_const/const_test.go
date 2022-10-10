package const_test

import (
	"fmt"
	"io"
	"testing"
)

// 连续位常量, 表状态
const (
	Readable   = 1 << iota // 0001
	Writeable              // 0010
	Executable             // 0100
)

func TestConstant(t *testing.T) {
	a := 1 // 0111
	// a = a &^ Readable   // 清零
	t.Log(a&Readable, a&Writeable, a&Executable)
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}

// BenchmarkConst 如果一个值创建后不会变动，定义为常量!
// 且 常量 比 变量易读。
func BenchmarkConst(b *testing.B) {
	b.Run("var", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			func() {
				str := "Hello world"
				fmt.Fprint(io.Discard, str)
			}()
		}
	})
	b.Run("const", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			func() {
				const str = "Hello world"
				fmt.Fprint(io.Discard, str)
			}()
		}
	})

}
