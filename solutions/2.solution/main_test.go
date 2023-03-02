package solutions

import (
	"fmt"
	"testing"

	"github.com/neverovski/leetcode-go/structures"

	"github.com/stretchr/testify/assert"
)

type param struct {
	l1 []int
	l2 []int
}

type testCase struct {
	in  param
	out []int
}

var cases = []testCase{
	{
		in:  param{[]int{}, []int{}},
		out: []int{},
	},
	{
		in:  param{[]int{2, 4, 3}, []int{5, 6, 4}},
		out: []int{7, 0, 8},
	},
	{
		in:  param{[]int{1}, []int{1}},
		out: []int{2},
	},
	{
		in:  param{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}},
		out: []int{8, 9, 9, 9, 0, 0, 0, 1},
	},
}

func TestAddTwoNumbers(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%v,%v", tt.in, tt.out), func(t *testing.T) {
			l1 := structures.TransformIntToList(tt.in.l1)
			l2 := structures.TransformIntToList(tt.in.l2)
			l3 := structures.TransformListToInt(addTwoNumbers(l1, l2))

			assert.Equal(t, tt.out, l3)
		})
	}
}
