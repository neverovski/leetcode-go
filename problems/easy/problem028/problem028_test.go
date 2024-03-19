package problem028

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	haystack string
	needle   string
	out      int
}

func TestStrStr(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{haystack: "sadbutsad", needle: "sad", out: 0},
		{haystack: "leetcode", needle: "leeto", out: -1},
		{haystack: "hello", needle: "ll", out: 2},
		{haystack: "a", needle: "a", out: 0},
		{haystack: "hello", needle: "world", out: -1},
		{haystack: "hello", needle: "", out: 0},
		{haystack: "", needle: "hello", out: -1},
		{haystack: "", needle: "", out: 0},
		{haystack: "mississippi", needle: "sip", out: 6},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, strStr(tc.haystack, tc.needle), fmt.Sprintf("TestStrStr number # %d", key+1))
	}
}

func TestStrStrUsingStrings(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{haystack: "sadbutsad", needle: "sad", out: 0},
		{haystack: "leetcode", needle: "leeto", out: -1},
		{haystack: "hello", needle: "ll", out: 2},
		{haystack: "a", needle: "a", out: 0},
		{haystack: "hello", needle: "world", out: -1},
		{haystack: "hello", needle: "", out: 0},
		{haystack: "", needle: "hello", out: -1},
		{haystack: "", needle: "", out: 0},
		{haystack: "mississippi", needle: "sip", out: 6},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, strStrUsingStrings(tc.haystack, tc.needle), fmt.Sprintf("TestStrStrUsingStrings number # %d", key+1))
	}
}

func BenchmarkStrStr(b *testing.B) {
	haystack := "hello"
	needle := "ll"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strStr(haystack, needle)
	}
}

func BenchmarkStrStrUsingStrings(b *testing.B) {
	haystack := "hello"
	needle := "ll"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strStrUsingStrings(haystack, needle)
	}
}
