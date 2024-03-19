package problem027

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums []int
	val  int
	out  int
}

func TestRemoveElement(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{nums: []int{3, 2, 2, 3}, val: 3, out: 2},
		{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, out: 5},
		{nums: []int{2}, val: 2, out: 0},
		{nums: []int{1, 1, 1, 2, 2, 2}, val: 2, out: 3},
		{nums: []int{}, val: 2, out: 0},
		{nums: []int{1, 2, 3, 4, 5}, val: 6, out: 5},
		{nums: []int{1, 2}, val: 2, out: 1},
		{nums: []int{2, 1}, val: 2, out: 1},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, removeElement(tc.nums, tc.val), fmt.Sprintf("TestRemoveElement number # %d", key+1))
	}
}

func BenchmarkRemoveElement(b *testing.B) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 3

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		removeElement(nums, val)
	}
}
