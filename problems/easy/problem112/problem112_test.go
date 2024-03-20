package problem112

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	root      *TreeNode
	targetSum int
	out       bool
}

func TestHasPathSum(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val:  8,
					Left: &TreeNode{Val: 13},
					Right: &TreeNode{
						Val:  4,
						Left: &TreeNode{Val: 1},
					},
				},
			},
			targetSum: 22,
			out:       true,
		},
		{
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			targetSum: 5,
			out:       false,
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, hasPathSum(tc.root, tc.targetSum), fmt.Sprintf("TestHasPathSum number # %d", key+1))
	}
}

func BenchmarkHasPathSum(b *testing.B) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val:   4,
				Right: &TreeNode{Val: 1},
			},
		},
	}
	targetSum := 22

	for i := 0; i < b.N; i++ {
		hasPathSum(root, targetSum)
	}
}

func TestHasPathSumUsingStack(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val:  8,
					Left: &TreeNode{Val: 13},
					Right: &TreeNode{
						Val:  4,
						Left: &TreeNode{Val: 1},
					},
				},
			},
			targetSum: 22,
			out:       true,
		},
		{
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			targetSum: 5,
			out:       false,
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, hasPathSumUsingStack(tc.root, tc.targetSum), fmt.Sprintf("TestHasPathSumUsingStack number # %d", key+1))
	}
}

func BenchmarkHasPathSumUsingStack(b *testing.B) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val:   4,
				Right: &TreeNode{Val: 1},
			},
		},
	}
	targetSum := 22

	for i := 0; i < b.N; i++ {
		hasPathSumUsingStack(root, targetSum)
	}
}
