package problem088

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
	out   []int
}

func TestMerge(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			out:   []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			out:   []int{1},
		},
		{
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			out:   []int{1},
		},
		{
			nums1: []int{1, 1, 1, 0, 0, 0},
			m:     3,
			nums2: []int{1, 1, 1},
			n:     3,
			out:   []int{1, 1, 1, 1, 1, 1},
		},
		{
			nums1: []int{1, 5, 7, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			out:   []int{1, 2, 5, 5, 6, 7},
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, merge(tc.nums1, tc.m, tc.nums2, tc.n), fmt.Sprintf("TestMerge number # %d", key+1))
	}
}

func TestMergeUsingSort(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			out:   []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			out:   []int{1},
		},
		{
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			out:   []int{1},
		},
		{
			nums1: []int{1, 1, 1, 0, 0, 0},
			m:     3,
			nums2: []int{1, 1, 1},
			n:     3,
			out:   []int{1, 1, 1, 1, 1, 1},
		},
		{
			nums1: []int{1, 5, 7, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			out:   []int{1, 2, 5, 5, 6, 7},
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, mergeUsingSort(tc.nums1, tc.m, tc.nums2, tc.n), fmt.Sprintf("TestMergeUsingSort number # %d", key+1))
	}
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
