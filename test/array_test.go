package test

import (
	"testing"
)

func TestArray(t *testing.T) {
	array()
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array()
	}
}

func array() {
	a1 := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		a1[i] = i
	}

	b1 := make([]int, 10000)
	copy(b1, a1)

	b1 = append(b1[:1], b1[2:]...)
	b1 = append(b1[:100], b1[101:]...)
	b1 = append(b1[:500], b1[501:]...)
	b1 = append(b1[:1000], b1[1001:]...)
	b1 = append(b1[:5000], b1[5001:]...)

	for _, i := range a1 {
		bool := false
		for idx, j := range b1 {
			if i == j {
				bool = true
				b1 = append(b1[0:idx], b1[idx+1:]...)
			}
		}
		if !bool {
			// fmt.Println(i)
		}
	}
	// fmt.Println(len(b1))
}
