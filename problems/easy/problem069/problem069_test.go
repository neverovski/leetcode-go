package problem069

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	a := assert.New(t)

	a.Equal(1, mySqrt(1))
	a.Equal(0, mySqrt(0))
	a.Equal(2, mySqrt(8))
	a.Equal(2, mySqrt(4))
	a.Equal(3, mySqrt(9))
	a.Equal(15, mySqrt(225))
}
