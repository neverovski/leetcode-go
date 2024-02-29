package solution_027

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	a := assert.New(t)

	a.Equal(2, RemoveElement([]int{3, 2, 2, 3}, 3))
	a.Equal(5, RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	a.Equal(0, RemoveElement([]int{2}, 2))
}
