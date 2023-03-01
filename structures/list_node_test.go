package structures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	l    *ListNode
	nums []int
}

var cases = []testCase{
	{
		l: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
		nums: []int{1, 2, 4},
	},
	{
		l:    &ListNode{},
		nums: []int{0},
	},
	{
		l:    nil,
		nums: []int{},
	},
	{
		l: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 7,
						},
					},
				},
			},
		},
		nums: []int{3, 4, 5, 6, 7},
	},
	{
		l: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
			},
		},
		nums: []int{0, 1},
	},
}

func TestTransformListToInt(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%v,%v", tt.l, tt.nums), func(t *testing.T) {
			assert.Equal(t, tt.nums, TransformListToInt(tt.l))
		})
	}
}

func TestTransformIntsToList(t *testing.T) {
	ast := assert.New(t)

	ast.Nil(TransformIntToList([]int{}))
	l := TransformIntToList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	fmt.Printf("l1: %+v\n", l)
	for i := 1; l != nil; i++ {
		ast.Equal(i, l.Val)
		l = l.Next
	}
}
