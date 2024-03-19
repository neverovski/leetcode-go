package problem069

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  int
	out int
}

func TestMySqrt(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{in: 1, out: 1},
		{in: 0, out: 0},
		{in: 8, out: 2},
		{in: 4, out: 2},
		{in: 9, out: 3},
		{in: 225, out: 15},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, mySqrt(tc.in), fmt.Sprintf("TestMySqrt number # %d", key+1))
	}
}
