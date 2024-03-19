package problem067

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	a   string
	b   string
	out string
}

func TestAddBinary(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{a: "11", b: "1", out: "100"},
		{a: "1011", b: "1011", out: "10110"},
		{a: "1010", b: "1011", out: "10101"},
		{a: "101", b: "1011", out: "10000"},
		{a: "1", b: "10000", out: "10001"},
		{a: "0", b: "0", out: "0"},
		{a: "0", b: "1", out: "1"},
		{a: "1", b: "1", out: "10"},
		{a: "11", b: "11", out: "110"},
		{a: "111", b: "1", out: "1000"},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, addBinary(tc.a, tc.b), fmt.Sprintf("TestAddBinary number # %d", key+1))
	}
}

func BenchmarkAddBinary(b *testing.B) {
	a := "1011"
	bin := "1010"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		addBinary(a, bin)
	}
}
