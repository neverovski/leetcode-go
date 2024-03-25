package problem121

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  []int
	out int
}

func TestMaxProfit(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in:  []int{7, 1, 5, 3, 6, 4},
			out: 5,
		},
		{
			in:  []int{7, 6, 4, 3, 1},
			out: 0,
		},
		{
			in:  []int{7, 5, 2, 3, 4},
			out: 2,
		},
		{
			in:  []int{7, 5, 3, 2, 1},
			out: 0,
		},
		{
			in:  []int{2, 5, 1},
			out: 3,
		},
		{
			in:  []int{2, 4, 1},
			out: 2,
		},
		{
			in:  []int{2, 4, 1, 3},
			out: 2,
		},
		{
			in:  []int{2, 4, 5},
			out: 3,
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, maxProfit(tc.in), fmt.Sprintf("TestMaxProfit number # %d", key+1))
	}
}
