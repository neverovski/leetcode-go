package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	nums   []int
	target int
}

type testCase struct {
	in  param
	out []int
}

var cases = []testCase{
	{
		in:  param{[]int{2, 7, 11, 15}, 9},
		out: []int{0, 1},
	},
	{
		in:  param{[]int{3, 2, 4}, 6},
		out: []int{1, 2},
	},
	{
		in:  param{[]int{3, 3}, 6},
		out: []int{0, 1},
	},
	{
		in:  param{[]int{3, 2, 4}, 5},
		out: []int{0, 1},
	},
	{
		in:  param{[]int{8, 7, 3, 3, 4, 2}, 11},
		out: []int{0, 2},
	},
	{
		in:  param{[]int{0, 1}, 1},
		out: []int{0, 1},
	},
	{
		in:  param{[]int{0, 3}, 5},
		out: []int{},
	},
}

func TestTwoSum(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%v,%v", tt.in, tt.out), func(t *testing.T) {
			assert.Equal(t, tt.out, twoSum(tt.in.nums, tt.in.target))
		})
	}
}
