package problem013

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	a := assert.New(t)

	a.Equal(3, romanToInt("III"))
	a.Equal(54, romanToInt("LIV"))
	a.Equal(58, romanToInt("LVIII"))
	a.Equal(1994, romanToInt("MCMXCIV"))
	a.Equal(1, romanToInt("I"))
	a.Equal(3999, romanToInt("MMMCMXCIX"))
	a.Equal(4, romanToInt("IV"))
	a.Equal(9, romanToInt("IX"))
	a.Equal(40, romanToInt("XL"))
	a.Equal(90, romanToInt("XC"))
	a.Equal(400, romanToInt("CD"))
	a.Equal(900, romanToInt("CM"))
}

func BenchmarkRomanToInt(b *testing.B) {
	str := "MCMXCIV"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		romanToInt(str)
	}
}
