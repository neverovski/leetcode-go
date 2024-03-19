package problem111

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDepth(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		in  *TreeNode
		out int
	}{
		{
			in: &TreeNode{
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
			out: 3,
		},
		{
			in: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
							Right: &TreeNode{
								Val: 6,
							},
						},
					},
				},
			},
			out: 5,
		},
		{
			in: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			out: 2,
		},
	}

	for key, tt := range tests {
		a.Equal(tt.out, minDepth(tt.in), fmt.Sprintf("Test case number # %d", key+1))
	}
}
