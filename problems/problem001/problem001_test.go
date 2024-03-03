package problem001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	a := assert.New(t)

	a.ElementsMatch([]int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	a.ElementsMatch([]int{1, 2}, twoSum([]int{3, 2, 4}, 6))
	a.ElementsMatch([]int{0, 1}, twoSum([]int{3, 3}, 6))
	a.ElementsMatch([]int{1, 2}, twoSum([]int{1, 2, 3, 4}, 5))
	a.ElementsMatch([]int{2, 4}, twoSum([]int{1, 1, 2, 3, 4}, 6))
	a.ElementsMatch([]int{1, 2}, twoSum([]int{-1, -2, -3, -4}, -5))
	a.ElementsMatch([]int{0, 1}, twoSum([]int{0, 0, 0}, 0))
	a.ElementsMatch([]int{}, twoSum([]int{1, 2, 3, 4}, 10))
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}
