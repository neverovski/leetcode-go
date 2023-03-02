package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	in  *TreeNode
	out bool
}

var cases = []testCase{
	{
		in: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 3},
			},
		},
		out: true,
	},
	{
		in: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 3},
			},
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
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 5},
			},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 3},
			},
		},
		out: false,
	},
}

func BenchmarkIsSymmetricLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isSymmetricLoop(cases[0].in)
	}
}

func BenchmarkIsSymmetricRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isSymmetricRecursive(cases[0].in)
	}
}

func TestIsSymmetricLoop(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%v,%v", tt.in, tt.out), func(t *testing.T) {
			assert.Equal(t, tt.out, isSymmetricLoop(tt.in))
		})
	}
}

func TestIsSymmetricRecursive(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%v,%v", tt.in, tt.out), func(t *testing.T) {
			assert.Equal(t, tt.out, isSymmetricRecursive(tt.in))
		})
	}
}
