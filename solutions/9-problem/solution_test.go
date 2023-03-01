package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	in  int
	out bool
}

var cases = []testCase{
	{121, true},
	{-121, false},
	{10, false},
	{123, false},
	{1231, false},
	{1331, true},
	{0, true},
	{1, true},
	{1234321, true},
}

func TestIsPalindrome(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%v,%v", tt.in, tt.out), func(t *testing.T) {
			assert.Equal(t, tt.out, isPalindrome(tt.in))
		})
	}
}
