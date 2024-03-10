package problem100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T) {
	a := assert.New(t)

	a.Equal(true, isSameTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}))
	a.Equal(false, isSameTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, &TreeNode{Val: 1, Right: &TreeNode{Val: 3}}))
	a.Equal(false, isSameTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}}, &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}))
	a.Equal(false, isSameTree(nil, &TreeNode{}))
}

func TestIsSameTreeUsingRecursive(t *testing.T) {
	a := assert.New(t)

	a.Equal(true, isSameTreeUsingRecursive(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}))
	a.Equal(false, isSameTreeUsingRecursive(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, &TreeNode{Val: 1, Right: &TreeNode{Val: 3}}))
	a.Equal(false, isSameTreeUsingRecursive(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}}, &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}))
	a.Equal(false, isSameTreeUsingRecursive(nil, &TreeNode{}))
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
