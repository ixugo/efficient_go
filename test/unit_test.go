package unit_test

import (
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
func add(a int) int {
	a++
	return a
}
