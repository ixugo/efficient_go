package sort_test

import "testing"

// insertion1 插入排序
// 此方法效率较低
func insertion1(nums []int) {
	n := len(nums)

	// 0,n-1
	// i+1, 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j > 0; j-- {
			n1 := nums[j]
			n2 := nums[j-1]
			if n1 < n2 {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

// insertion2 插入排序
func insertion2(nums []int) {
	n := len(nums)

	// 0,n-1
	// i+1,0
	for i := 0; i < n-1; i++ {
		j := i + 1
		tmp := nums[j]
		for j > 0 && nums[j-1] > tmp {
			nums[j] = nums[j-1]
			j--
		}
		nums[j] = tmp
	}
}

func TestInsertion1(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{9, 1, 0, 2, 4, 1, 5, 8, 7, 3, 54, 12, 59, 39}
		insertion1(nums)
		t.Log(nums)
	})
	t.Run("2", func(t *testing.T) {
		nums := []int{9, 1, 0, 2, 4, 1, 5, 8, 7, 3, 54, 12, 59, 39}
		insertion2(nums)
		t.Log(nums)
	})
}

func BenchmarkInsertion(b *testing.B) {
	b.Run("1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			nums := []int{9, 1, 0, 2, 4, 1, 5, 8, 7, 3, 54, 12, 59, 39}
			insertion1(nums)
		}
	})

	b.Run("2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			nums := []int{9, 1, 0, 2, 4, 1, 5, 8, 7, 3, 54, 12, 59, 39}
			insertion2(nums)
		}
	})

}
