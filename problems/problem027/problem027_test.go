package problem027

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	a := assert.New(t)

	a.Equal(2, removeElement([]int{3, 2, 2, 3}, 3))
	a.Equal(5, removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	a.Equal(0, removeElement([]int{2}, 2))
}
