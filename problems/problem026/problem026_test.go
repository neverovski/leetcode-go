package problem026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	a := assert.New(t)

	a.Equal(2, removeDuplicates([]int{1, 1, 2}))
	a.Equal(5, removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	a.Equal(1, removeDuplicates([]int{1, 1, 1, 1, 1}))
	a.Equal(3, removeDuplicates([]int{1, 2, 3}))
	a.Equal(0, removeDuplicates([]int{}))
	a.Equal(4, removeDuplicates([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}))
}
