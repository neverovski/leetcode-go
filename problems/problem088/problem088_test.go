package problem088

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	a := assert.New(t)

	a.Equal([]int{1, 2, 2, 3, 5, 6}, merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
	a.Equal([]int{1}, merge([]int{1}, 1, []int{}, 0))
	a.Equal([]int{1}, merge([]int{0}, 0, []int{1}, 1))
	a.Equal([]int{1, 1, 1, 1, 1, 1}, merge([]int{1, 1, 1, 0, 0, 0}, 3, []int{1, 1, 1}, 3))
	a.Equal([]int{1, 2, 5, 5, 6, 7}, merge([]int{1, 5, 7, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
}

func TestMergeUsingSort(t *testing.T) {
	a := assert.New(t)

	a.Equal([]int{1, 2, 2, 3, 5, 6}, mergeUsingSort([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
	a.Equal([]int{1}, mergeUsingSort([]int{1}, 1, []int{}, 0))
	a.Equal([]int{1}, mergeUsingSort([]int{0}, 0, []int{1}, 1))
	a.Equal([]int{1, 1, 1, 1, 1, 1}, mergeUsingSort([]int{1, 1, 1, 0, 0, 0}, 3, []int{1, 1, 1}, 3))
	a.Equal([]int{1, 2, 5, 5, 6, 7}, mergeUsingSort([]int{1, 5, 7, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
}

func BenchmarkMerge(b *testing.B) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		merge(nums1, m, nums2, n)
	}
}

func BenchmarkMergeUsingSort(b *testing.B) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mergeUsingSort(nums1, m, nums2, n)
	}
}
