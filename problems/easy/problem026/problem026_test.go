package problem026

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  []int
	out int
}

func TestRemoveDuplicates(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{in: []int{1, 1, 2}, out: 2},
		{in: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, out: 5},
		{in: []int{1, 1, 1, 1, 1}, out: 1},
		{in: []int{1, 2, 3}, out: 3},
		{in: []int{}, out: 0},
		{in: []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}, out: 4},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, removeDuplicates(tc.in), fmt.Sprintf("TestRemoveDuplicates number # %d", key+1))
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	nums := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		removeDuplicates(nums)
	}
}
