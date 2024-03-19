package problem001

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums   []int
	target int
	out    []int
}

func TestTwoSum(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{nums: []int{2, 7, 11, 15}, target: 9, out: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, out: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, out: []int{0, 1}},
		{nums: []int{1, 2, 3, 4}, target: 5, out: []int{1, 2}},
		{nums: []int{1, 1, 2, 3, 4}, target: 6, out: []int{2, 4}},
		{nums: []int{-1, -2, -3, -4}, target: -5, out: []int{1, 2}},
		{nums: []int{0, 0, 0}, target: 0, out: []int{0, 1}},
		{nums: []int{1, 2, 3, 4}, target: 10, out: []int{}},
	}

	for key, tc := range testCases {
		a.ElementsMatch(tc.out, twoSum(tc.nums, tc.target), fmt.Sprintf("TestTwoSum number # %d", key+1))
	}
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}
