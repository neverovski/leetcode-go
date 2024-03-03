package problem014

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	a := assert.New(t)

	a.Equal("fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	a.Equal("", longestCommonPrefix([]string{"dog", "racecar", "car"}))
	a.Equal("a", longestCommonPrefix([]string{"abc", "a", "ab"}), "one string is a prefix of the others")
	a.Equal("", longestCommonPrefix([]string{"", "b", "c"}), "one string is empty")
	a.Equal("abc", longestCommonPrefix([]string{"abc", "abc", "abc"}), "all strings are the same")
	a.Equal("", longestCommonPrefix([]string{"abc", "def", "ghi"}), "no common prefix")
	a.Equal("", longestCommonPrefix([]string{}), "empty slice")
}
