package problem020

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  string
	out bool
}

func TestIsValid(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{in: "()", out: true},
		{in: "()[]{}", out: true},
		{in: "(]", out: false},
		{in: "([)]", out: false},
		{in: "{[]}", out: true},
		{in: "", out: false},
		{in: "[]{}[]", out: true},
		{in: "[()]", out: true},
		{in: "[)", out: false},
		{in: "[", out: false},
		{in: "[[", out: false},
		{in: ")(){}", out: false},
		{in: "(])", out: false},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, isValid(tc.in), fmt.Sprintf("TestIsValid number # %d", key+1))
	}
}

func BenchmarkIsValid(b *testing.B) {
	str := "[]{}[]"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		isValid(str)
	}
}
