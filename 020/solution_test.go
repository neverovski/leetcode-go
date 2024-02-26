package solution_020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	a := assert.New(t)

	a.Equal(true, IsValid("[]"))
	a.Equal(true, IsValid("[]{}[]"))
	a.Equal(false, IsValid("[)"))
	a.Equal(true, IsValid("[()]"))
	a.Equal(false, IsValid("["))
	a.Equal(false, IsValid("[["))
	a.Equal(false, IsValid("]"))
	a.Equal(false, IsValid(")(){}"))
	a.Equal(false, IsValid("(])"))
}
