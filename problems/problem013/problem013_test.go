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
}
