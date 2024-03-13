package problem108

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedArrayToBST(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		input  []int
		output *TreeNode
	}{
		{
			input: []int{-10, -3, 0, 5, 9},
			output: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:  -3,
					Left: &TreeNode{Val: -10},
				},
				Right: &TreeNode{
					Val:  9,
					Left: &TreeNode{Val: 5},
				},
			},
		},
		{
			input: []int{1, 3},
			output: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 1},
			},
		},
		{
			input: []int{0, 1, 2, 3, 4, 5, 6},
			output: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 6},
				},
			},
		},
	}

	for _, tt := range tests {
		a.Equal(tt.output, sortedArrayToBST(tt.input))
	}
}
