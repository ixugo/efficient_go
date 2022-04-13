package const_test

import (
	"fmt"
	"io"
	"testing"
)

// BenchmarkConst 如果一个值创建后不会变动，定义为常量!
// 切 常量 比 变量易读。
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
