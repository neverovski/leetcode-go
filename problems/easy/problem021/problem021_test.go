package problem021

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	list1 *ListNode
	list2 *ListNode
	out   *ListNode
}

func TestMergeTwoLists(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			list1: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}},
			list2: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
			out:   &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: nil}}}}}},
		},
		{
			list1: nil,
			list2: &ListNode{Val: 0, Next: nil},
			out:   &ListNode{Val: 0, Next: nil},
		},
		{
			list1: nil,
			list2: nil,
			out:   nil,
		},
	}

	for key, tc := range testCases {
		a.Equal(tc.out, mergeTwoLists(tc.list1, tc.list2), fmt.Sprintf("TestMergeTwoLists number # %d", key+1))
	}
}

func BenchmarkMergeTwoLists(b *testing.B) {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mergeTwoLists(list1, list2)
	}
}
