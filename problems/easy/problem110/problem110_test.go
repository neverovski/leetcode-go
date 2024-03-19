package problem110

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  *TreeNode
	out bool
}

func TestIsBalanced(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			out: true,
		},
		{
			in: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{Val: 2},
			},
			out: false,
		},
		{
			in:  nil,
			out: true,
		},
		{
			in: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			out: false,
		},
		{
			in: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			out: true,
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, isBalanced(tc.in), fmt.Sprintf("TestIsBalanced number # %d", key+1))
	}
}
