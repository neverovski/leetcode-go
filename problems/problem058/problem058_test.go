package problem058

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	a := assert.New(t)

	a.Equal(5, lengthOfLastWord("Hello World"))
	a.Equal(4, lengthOfLastWord("   fly me   to   the moon  "))
	a.Equal(6, lengthOfLastWord("luffy is still joyboy"))
	a.Equal(1, lengthOfLastWord("a"))
}
