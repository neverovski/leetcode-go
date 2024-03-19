package problem009

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  int
	out bool
}

func TestIsPalindrome(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{in: 121, out: true},
		{in: 10, out: false},
		{in: -123, out: false},
		{in: 5, out: true},
		{in: 1221, out: true},
		{in: 1231, out: false},
		{in: 1, out: true},
		{in: -1, out: false},
		{in: 0, out: true},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, isPalindrome(tc.in), fmt.Sprintf("TestIsPalindrome number # %d", key+1))
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	num := 121

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		isPalindrome(num)
	}
}
