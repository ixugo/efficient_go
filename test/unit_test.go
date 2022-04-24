package test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// 表格测试
func TestTable(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{2, 3, 4}

	for i, v := range inputs {
		require.EqualValues(t, add(v), expected[i])
	}
}

// 性能测试
func BenchmarkConcat1(b *testing.B) {
	e := [...]string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := ""
		for _, v := range e {
			result += v
		}
	}
	b.StopTimer()
}
func BenchmarkConcat2(b *testing.B) {
	e := [...]string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result bytes.Buffer
		for _, v := range e {
			result.WriteString(v)
		}
	}
	b.StopTimer()
}

func BenchmarkConcat3(b *testing.B) {
	e := [...]string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result strings.Builder
		for _, v := range e {
			result.WriteString(v)
		}
	}
	b.StopTimer()
}
