package problem009

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	a := assert.New(t)

	a.Equal(true, isPalindrome(121))
	a.Equal(false, isPalindrome(10))
	a.Equal(false, isPalindrome(-123))
	a.Equal(true, isPalindrome(5))
	a.Equal(true, isPalindrome(1221), "even number of digits")
	a.Equal(false, isPalindrome(1231), "not a palindrome")
	a.Equal(true, isPalindrome(1), "single digit")
	a.Equal(false, isPalindrome(-1), "negative single digit")
	a.Equal(true, isPalindrome(0), "zero")
}

func BenchmarkIsPalindrome(b *testing.B) {
	num := 121

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		isPalindrome(num)
	}
}
