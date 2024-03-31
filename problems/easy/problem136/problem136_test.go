package problem136

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  []int
	out int
}

func TestSingleNumber(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in:  []int{2, 2, 1},
			out: 1,
		},
		{
			in:  []int{4, 1, 2, 1, 2},
			out: 4,
		},
		{
			in:  []int{1},
			out: 1,
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, singleNumber(tc.in), fmt.Sprintf("TestSingleNumber number # %d", key+1))
	}
}

func TestSingleNumberUsingBit(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in:  []int{2, 2, 1},
			out: 1,
		},
		{
			in:  []int{4, 1, 2, 1, 2},
			out: 4,
		},
		{
			in:  []int{1},
			out: 1,
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, singleNumberUsingBit(tc.in), fmt.Sprintf("TestSingleNumberUsingBit number # %d", key+1))
	}
}

func BenchmarkSingleNumber(b *testing.B) {
	nums := []int{4, 1, 2, 1, 2}
	for i := 0; i < b.N; i++ {
		singleNumber(nums)
	}
}

func BenchmarkSingleNumberUsingBit(b *testing.B) {
	nums := []int{4, 1, 2, 1, 2}
	for i := 0; i < b.N; i++ {
		singleNumberUsingBit(nums)
	}
}
