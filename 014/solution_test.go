package solution_014

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	a := assert.New(t)

	a.Equal("fl", LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	a.Equal("", LongestCommonPrefix([]string{"dog", "racecar", "car"}))
}
