package problem070

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  int
	out int
}

func TestClimbStairs(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{in: 1, out: 1},
		{in: 2, out: 2},
		{in: 3, out: 3},
		{in: 4, out: 5},
		{in: 5, out: 8},
		{in: 6, out: 13},
		{in: 7, out: 21},
		{in: 8, out: 34},
		{in: 9, out: 55},
		{in: 10, out: 89},
		{in: 11, out: 144},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, climbStairs(tc.in), fmt.Sprintf("TestClimbStairs number # %d", key+1))
	}
}

func BenchmarkClimbStairs(b *testing.B) {
	n := 30

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		climbStairs(n)
	}
}
