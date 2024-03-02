package solution_028

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	a := assert.New(t)

	a.Equal(0, strStr("sadbutsad", "sad"))
	a.Equal(-1, strStr("leetcode", "leeto"))
	a.Equal(2, strStr("hello", "ll"))
	a.Equal(0, strStr("a", "a"))
	a.Equal(-1, strStr("hello", "world"))
	a.Equal(0, strStr("hello", ""))
	a.Equal(-1, strStr("", "hello"))
	a.Equal(0, strStr("", ""))
	a.Equal(6, strStr("mississippi", "sip"))
}

func TestStrStrUsingStrings(t *testing.T) {
	a := assert.New(t)

	a.Equal(0, strStrUsingStrings("sadbutsad", "sad"))
	a.Equal(-1, strStrUsingStrings("leetcode", "leeto"))
	a.Equal(2, strStrUsingStrings("hello", "ll"))
	a.Equal(0, strStrUsingStrings("a", "a"))
	a.Equal(-1, strStrUsingStrings("hello", "world"))
	a.Equal(0, strStrUsingStrings("hello", ""))
	a.Equal(-1, strStrUsingStrings("", "hello"))
	a.Equal(0, strStrUsingStrings("", ""))
	a.Equal(6, strStrUsingStrings("mississippi", "sip"))
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
