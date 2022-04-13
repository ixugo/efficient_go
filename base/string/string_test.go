package string

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

const (
	COUNT = 1000
)

func TestForRangeStr(t *testing.T) {
	s := "世界 Hello world"

	var buf [4]byte
	for i, v := range s {
		rl := utf8.RuneLen(v)
		si := i + rl
		// 字符串是不可变的，所以只能作为来源
		// 数组也是不可变的，可以通过转为切片，来操作底层数组。
		copy(buf[:], s[i:si])
		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, v, v, buf[:rl])
	}
}

// 通过比较发现
// builder > buffer > add > sprintf

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ""
		for i := 0; i < COUNT; i++ {
			s = fmt.Sprintf("%s%s", s, "string")
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ""
		for i := 0; i < COUNT; i++ {
			s += "string"
		}
	}
}

func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		for i := 0; i < COUNT; i++ {
			buffer.WriteString("string")
		}
	}
}
func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for i := 0; i < COUNT; i++ {
			builder.WriteString("string")
		}
	}
}
