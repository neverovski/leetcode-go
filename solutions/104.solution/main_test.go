package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	in  *TreeNode
	out int
}

var cases = []testCase{
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
		out: 3,
	},
	{
		in: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		out: 2,
	},
	{
		in: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   3,
				Right: &TreeNode{Val: 5},
			},
		},
		out: 3,
	},
}

func BenchmarkMaxDepthLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxDepthLoop(cases[0].in)
	}
}

func BenchmarkMaxDepthRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxDepthRecursive(cases[0].in)
	}
}

func TestMaxDepthLoop(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%v,%v", tt.in, tt.out), func(t *testing.T) {
			assert.Equal(t, tt.out, maxDepthLoop(tt.in))
		})
	}
}

func TestMaxDepthRecursive(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%v,%v", tt.in, tt.out), func(t *testing.T) {
			assert.Equal(t, tt.out, maxDepthRecursive(tt.in))
		})
	}
}
