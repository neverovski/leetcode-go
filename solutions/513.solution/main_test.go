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
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		out: 1,
	},
	{
		in: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
		},
		out: 7,
	},
	{
		in: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
		},
		out: 4,
	},
	{
		in: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 6,
				},
			},
		},
		out: 8,
	},
	{
		in: &TreeNode{
			Val: 0,
			Right: &TreeNode{
				Val: -1,
			},
		},
		out: -1,
	},
	{
		in: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: -1,
			},
		},
		out: -1,
	},
}

func BenchmarkFindBottomLeftValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findBottomLeftValue(cases[0].in)
	}
}

func TestFindBottomLeftValue(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%v,%v", tt.in, tt.out), func(t *testing.T) {
			assert.Equal(t, tt.out, findBottomLeftValue(tt.in))
		})
	}
}
