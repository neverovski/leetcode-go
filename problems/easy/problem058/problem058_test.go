package problem058

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  string
	out int
}

func TestLengthOfLastWord(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{in: "Hello World", out: 5},
		{in: "   fly me   to   the moon  ", out: 4},
		{in: "luffy is still joyboy", out: 6},
		{in: "a", out: 1},
		{in: "", out: 0},
		{in: "   ", out: 0},
		{in: "Hello", out: 5},
		{in: "   Hello   ", out: 5},
		{in: "Hello   World   ", out: 5},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, lengthOfLastWord(tc.in), fmt.Sprintf("TestLengthOfLastWord number # %d", key+1))
	}
}

func BenchmarkLengthOfLastWord(b *testing.B) {
	str := "luffy is still joyboy"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lengthOfLastWord(str)
	}
}
