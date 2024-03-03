package problem020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	a := assert.New(t)

	a.Equal(true, isValid("[]"))
	a.Equal(true, isValid("[]{}[]"))
	a.Equal(false, isValid("[)"))
	a.Equal(true, isValid("[()]"))
	a.Equal(false, isValid("["))
	a.Equal(false, isValid("[["))
	a.Equal(false, isValid("]"))
	a.Equal(false, isValid(")(){}"))
	a.Equal(false, isValid("(])"))
}
