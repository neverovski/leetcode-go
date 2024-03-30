package problem125

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  string
	out bool
}

func TestIsPalindrome(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in:  "A man, a plan, a canal: Panama",
			out: true,
		},
		{
			in:  "race a car",
			out: false,
		},
		{
			in:  " ",
			out: true,
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, isPalindrome(tc.in), fmt.Sprintf("TestIsPalindrome number # %d", key+1))
	}
}
