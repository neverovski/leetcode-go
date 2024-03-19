package problem035

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums   []int
	target int
	out    int
}

func TestSearchInsert(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{nums: []int{1, 3, 5, 6}, target: 5, out: 2},
		{nums: []int{1, 3, 5, 6}, target: 2, out: 1},
		{nums: []int{1, 3, 5, 6}, target: 7, out: 4},
		{nums: []int{1, 3, 5, 6}, target: 0, out: 0},
		{nums: []int{1, 3, 5, 6}, target: 6, out: 3},
		{nums: []int{}, target: 0, out: 0},
		{nums: []int{1, 3}, target: 3, out: 1},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, searchInsert(tc.nums, tc.target), fmt.Sprintf("TestSearchInsert number # %d", key+1))
	}
}

func BenchmarkSearchInsert(b *testing.B) {
	nums := []int{1, 3, 5, 6}
	target := 5

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		searchInsert(nums, target)
	}
}
