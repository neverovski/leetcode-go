package solution_009

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	a := assert.New(t)

	a.Equal(true, IsPalindrome(121))
	a.Equal(false, IsPalindrome(10))
	a.Equal(false, IsPalindrome(-123))
	a.Equal(true, IsPalindrome(5))
}
