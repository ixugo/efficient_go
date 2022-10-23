package slice

import (
	"fmt"
	"strconv"
	"testing"
)

// TestAppend 输出扩容规则
// 官方代码: https://github.com/golang/go/blob/master/src/runtime/slice.go#L181
// Go 1.18 扩容规则，1.18 以前不一样。
// 容量翻倍 < 所需容量，预估容量 = 所需容量
//
//	否则，< 256，双倍扩容
//	否则，2 倍扩容到 1.25 倍扩容平滑过渡
//
// 最后，匹配内存规格 8 * (2 * x)，x 从 0 递增
func TestAppend(t *testing.T) {
	var (
		data    []string
		lastCap int
	)
	for i := 0; i < 1e7; i++ {
		data = append(data, strconv.Itoa(i))
		if newCap := cap(data); lastCap != newCap {
			capChg := float64(newCap-lastCap) / float64(lastCap) * 100
			lastCap = newCap
			fmt.Printf("index[%-8d]\tcap[%-8d - %.2f]\n", i, newCap, capChg)
		}
	}
}

// TestZero int类型 除以零会发生 panic
// float 不会发生错误，会得到 NaN
func TestZero(t *testing.T) {
	a := 0.0
	b := 0.0
	fmt.Println(a / b)
}
