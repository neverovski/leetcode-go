package problem169

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  []int
	out int
}

func TestMajorityElement(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in:  []int{3, 2, 3},
			out: 3,
		},
		{
			in:  []int{2, 2, 1, 1, 1, 2, 2},
			out: 2,
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, majorityElement(tc.in), fmt.Sprintf("TestMajorityElement number # %d", key+1))
	}
}
