package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	in  *TreeNode
	out []int
}

var cases = []testCase{
	{
		in: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			}},
		out: []int{1, 3, 2},
	},
	{
		in:  nil,
		out: []int{},
	},
	{
		in:  &TreeNode{Val: 1},
		out: []int{1},
	},
}

func TestInorderTraversalUsingLoop(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%v,%v", tt.in, tt.out), func(t *testing.T) {
			assert.Equal(t, tt.out, inorderTraversalUsingLoop(tt.in))
		})
	}
}

func TestInorderTraversalRecursive(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%v,%v", tt.in, tt.out), func(t *testing.T) {
			assert.Equal(t, tt.out, inorderTraversalRecursive(tt.in))
		})
	}
}
