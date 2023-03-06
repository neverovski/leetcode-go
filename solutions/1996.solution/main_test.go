package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	in  [][]int
	out int
}

var cases = []testCase{
	{
		[][]int{{5, 5}, {6, 3}, {3, 6}},
		0,
	},
	{
		[][]int{{2, 2}, {3, 3}},
		1,
	},
	{
		[][]int{{1, 5}, {10, 4}, {4, 3}},
		1,
	},
	{
		[][]int{{7, 7}, {1, 2}, {9, 7}, {7, 3}, {3, 10}, {9, 8}, {8, 10}, {4, 3}, {1, 5}, {1, 5}},
		6,
	},
	{
		[][]int{{5, 5}, {6, 6}, {3, 6}, {7, 7}},
		3,
	},
	{
		[][]int{{7, 9}, {10, 7}, {6, 9}, {10, 4}, {7, 5}, {7, 10}},
		2,
	},
}

func BenchmarkSortSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortSlice(cases[0].in)
	}
}

func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Max(5, 6)
	}
}

func BenchmarkMathMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MathMax(5, 6)
	}
}

func TestNumberOfWeakCharacters(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%v,%v", tt.in, tt.out), func(t *testing.T) {
			assert.Equal(t, tt.out, numberOfWeakCharacters(tt.in))
		})
	}
}
