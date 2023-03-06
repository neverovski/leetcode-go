package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	in  string
	out string
}

var cases = []testCase{
	{"the sky is blue", "blue is sky the"},
	{"  hello world  ", "world hello"},
	{"a good   example", "example good a"},
	{"a good   example  ", "example good a"},
	{"   a  good   example   ", "example good a"},
}

func BenchmarkReverseWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseWords(cases[0].in)
	}
}

func TestReverseWords(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%v,%v", tt.in, tt.out), func(t *testing.T) {
			assert.Equal(t, tt.out, reverseWords(tt.in))
		})
	}
}
