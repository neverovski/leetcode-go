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
}
