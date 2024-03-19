package problem066

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  []int
	out []int
}

func TestPlusOne(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{in: []int{1, 2, 3}, out: []int{1, 2, 4}},
		{in: []int{4, 3, 2, 1}, out: []int{4, 3, 2, 2}},
		{in: []int{9}, out: []int{1, 0}},
		{in: []int{9, 9}, out: []int{1, 0, 0}},
		{in: []int{9, 9, 9}, out: []int{1, 0, 0, 0}},
		{in: []int{8, 9, 9, 9}, out: []int{9, 0, 0, 0}},
		{in: []int{7, 8, 9, 9}, out: []int{7, 9, 0, 0}},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, plusOne(tc.in), fmt.Sprintf("TestPlusOne number # %d", key+1))
	}
}

func BenchmarkPlusOne(b *testing.B) {
	digits := []int{1, 2, 3}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		plusOne(digits)
	}
}
