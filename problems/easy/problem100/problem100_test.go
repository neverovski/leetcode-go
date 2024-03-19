package problem100

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	p   *TreeNode
	q   *TreeNode
	out bool
}

func TestIsSameTree(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			p:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			q:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			out: true,
		},
		{
			p:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}},
			q:   &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
			out: false,
		},
		{
			p:   nil,
			q:   &TreeNode{},
			out: false,
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, isSameTree(tc.p, tc.q), fmt.Sprintf("TestIsSameTree number # %d", key+1))
	}
}

func TestIsSameTreeUsingRecursive(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			p:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			q:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			out: true,
		},
		{
			p:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}},
			q:   &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
			out: false,
		},
		{
			p:   nil,
			q:   &TreeNode{},
			out: false,
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, isSameTreeUsingRecursive(tc.p, tc.q), fmt.Sprintf("TestIsSameTreeUsingRecursive number # %d", key+1))
	}
}

func BenchmarkIsSameTree(b *testing.B) {
	tree1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	tree2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		isSameTree(tree1, tree2)
	}
}

func BenchmarkIsSameTreeUsingRecursive(b *testing.B) {
	tree1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	tree2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		isSameTreeUsingRecursive(tree1, tree2)
	}
}
