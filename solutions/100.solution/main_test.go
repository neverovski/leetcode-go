package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type inParam struct {
	p *TreeNode
	q *TreeNode
}

type testCase struct {
	in  *inParam
	out bool
}

var cases = []testCase{
	{
		in: &inParam{
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			q: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
		},
		out: true,
	},
	{
		in: &inParam{
			p: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			q: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
		},
		out: false,
	},
	{
		in: &inParam{
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 1},
			},
			q: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
		},
		out: false,
	},
	{
		in: &inParam{
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 3},
			},
			q: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 3},
			},
		},
		out: true,
	},
	{
		in: &inParam{
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 3},
			},
			q: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 3},
			},
		},
		out: false,
	},
	{
		in: &inParam{
			p: nil,
			q: nil,
		},
		out: true,
	},
	{
		in: &inParam{
			p: nil,
			q: &TreeNode{Val: 0},
		},
		out: false,
	},
}

func TestIsSameTree(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%v,%v", tt.in, tt.out), func(t *testing.T) {
			assert.Equal(t, tt.out, isSameTree(tt.in.p, tt.in.q))
		})
	}
}
