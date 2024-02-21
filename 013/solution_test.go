package solution_013

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	a := assert.New(t)

	a.Equal(3, RomanToInt("III"))
	a.Equal(54, RomanToInt("LIV"))
	a.Equal(58, RomanToInt("LVIII"))
	a.Equal(1994, RomanToInt("MCMXCIV"))
}
