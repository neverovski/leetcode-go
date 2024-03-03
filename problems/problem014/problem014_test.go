package problem014

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	a := assert.New(t)

	a.Equal("fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	a.Equal("", longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
