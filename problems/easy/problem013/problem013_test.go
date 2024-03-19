package problem013

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  string
	out int
}

func TestRomanToInt(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{in: "III", out: 3},
		{in: "LIV", out: 54},
		{in: "LVIII", out: 58},
		{in: "MCMXCIV", out: 1994},
		{in: "I", out: 1},
		{in: "MMMCMXCIX", out: 3999},
		{in: "IV", out: 4},
		{in: "IX", out: 9},
		{in: "XL", out: 40},
		{in: "XC", out: 90},
		{in: "CD", out: 400},
		{in: "CM", out: 900},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, romanToInt(tc.in), fmt.Sprintf("TestRomanToInt number # %d", key+1))
	}
}

func BenchmarkRomanToInt(b *testing.B) {
	str := "MCMXCIV"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		romanToInt(str)
	}
}
