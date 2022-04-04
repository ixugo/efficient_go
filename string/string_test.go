package string

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

const (
	COUNT = 1000
)

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
