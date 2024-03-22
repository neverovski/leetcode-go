package problem119

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  int
	out []int
}

func TestGetRow(t *testing.T) {
	a := assert.New(t)

	a.Equal([]int{1}, getRow(0))
	a.Equal([]int{1, 1}, getRow(1))
	a.Equal([]int{1, 3, 3, 1}, getRow(3))

	testCases := []TestCase{
		{
			in:  0,
			out: []int{1},
		},
		{
			in:  1,
			out: []int{1, 1},
		},
		{
			in:  3,
			out: []int{1, 3, 3, 1},
		},
		{
			in:  5,
			out: []int{1, 4, 6, 4, 1},
		},
		{
			in:  7,
			out: []int{1, 6, 15, 20, 15, 6, 1},
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, getRow(tc.in), fmt.Sprintf("TestGetRow number # %d", key+1))
	}
}
