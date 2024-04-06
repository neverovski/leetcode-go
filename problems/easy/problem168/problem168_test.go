package problem168

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  int
	out string
}

func TestConvertToTitle(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in:  1,
			out: "A",
		},
		{
			in:  28,
			out: "AB",
		},
		{
			in:  701,
			out: "ZY",
		},
		{
			in:  2147483647,
			out: "FXSHRXW",
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, convertToTitle(tc.in), fmt.Sprintf("TestConvertToTitle number # %d", key+1))
	}
}
