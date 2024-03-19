package problem014

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  []string
	out string
}

func TestLongestCommonPrefix(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{in: []string{"flower", "flow", "flight"}, out: "fl"},
		{in: []string{"dog", "racecar", "car"}, out: ""},
		{in: []string{"abc", "a", "ab"}, out: "a"},
		{in: []string{"", "b", "c"}, out: ""},
		{in: []string{"abc", "abc", "abc"}, out: "abc"},
		{in: []string{"abc", "def", "ghi"}, out: ""},
		{in: []string{}, out: ""},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, longestCommonPrefix(tc.in), fmt.Sprintf("TestLongestCommonPrefix number # %d", key+1))
	}
}

func BenchmarkLongestCommonPrefix(b *testing.B) {
	strs := []string{"flower", "flow", "flight"}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		longestCommonPrefix(strs)
	}
}
