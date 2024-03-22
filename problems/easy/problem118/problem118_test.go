package problem118

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  int
	out [][]int
}

func TestGenerate(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in:  1,
			out: [][]int{{1}},
		},
		{
			in:  5,
			out: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			in:  7,
			out: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}, {1, 6, 15, 20, 15, 6, 1}},
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, generate(tc.in), fmt.Sprintf("TestGenerate number # %d", key+1))
	}
}
