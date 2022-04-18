package allocationandgc

import "testing"

type Content struct {
	D [10000]int
}

// 避免内存分配和复制，可以减轻 gc 压力

// BenchmarkTransferObj 复杂对象尽量传递引用

func BenchmarkTransferObj(b *testing.B) {
	var arr [10000]Content
	// 值传递
	b.Run("value", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			withValue(arr)
		}
	})
	// 引用传递
	b.Run("refrence", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			withRefrence(&arr)
		}
	})
}

// BenchmarkSlice 使用切片，需要预分配合适内存
func BenchmarkSlice(b *testing.B) {
	b.Run("auto", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var data []int
			for i := 0; i < 10000; i++ {
				data = append(data, i)
			}
		}
	})

	b.Run("proper init", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data := make([]int, 0, 10000)
			for i := 0; i < 10000; i++ {
				data = append(data, i)
			}
		}
	})

}

func withValue(arr [10000]Content) int {
	return 0
}
func withRefrence(arr *[10000]Content) int {
	return 0
}
