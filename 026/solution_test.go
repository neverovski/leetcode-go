package solution_026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	a := assert.New(t)

	a.Equal(2, RemoveDuplicates([]int{1, 1, 2}))
	a.Equal(5, RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
