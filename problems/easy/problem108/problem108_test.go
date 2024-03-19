package problem108

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  []int
	out *TreeNode
}

func TestSortedArrayToBST(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in: []int{-10, -3, 0, 5, 9},
			out: &TreeNode{
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
			in: []int{1, 3},
			out: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 1},
			},
		},
		{
			in: []int{0, 1, 2, 3, 4, 5, 6},
			out: &TreeNode{
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

	for key, tc := range testCases {
		a.Equal(tc.out, sortedArrayToBST(tc.in), fmt.Sprintf("TestSortedArrayToBST number # %d", key+1))
	}
}
