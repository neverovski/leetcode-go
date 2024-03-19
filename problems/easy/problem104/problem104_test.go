package problem104

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  *TreeNode
	out int
}

func TestMaxDepth(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in:  nil,
			out: 0,
		},
		{
			in:  &TreeNode{Val: 1},
			out: 1,
		},
		{
			in:  &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			out: 2,
		},
		{
			in:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			out: 2,
		},
		{
			in:  &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			out: 3,
		},
		{
			in:  &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
			out: 3,
		},
		{
			in:  &TreeNode{Val: 0, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 6}}, Right: &TreeNode{Val: -1, Right: &TreeNode{Val: 8}}}},
			out: 4,
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, maxDepth(tc.in), fmt.Sprintf("TestMaxDepth number # %d", key+1))
	}
}

func BenchmarkMaxDepth(b *testing.B) {
	tree := &TreeNode{Val: 0, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 6}}, Right: &TreeNode{Val: -1, Right: &TreeNode{Val: 8}}}}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		maxDepth(tree)
	}
}
