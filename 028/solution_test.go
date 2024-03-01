package solution_028

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	a := assert.New(t)

	a.Equal(0, StrStr("sadbutsad", "sad"))
	a.Equal(-1, StrStr("leetcode", "leeto"))
	a.Equal(2, StrStr("hello", "ll"))
	a.Equal(0, StrStr("a", "a"))
}
