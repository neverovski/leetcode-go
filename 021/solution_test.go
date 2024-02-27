package solution_021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	a := assert.New(t)

	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	expected := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: nil}}}}}}
	a.Equal(expected, MergeTwoLists(list1, list2))

	list1 = nil
	list2 = &ListNode{Val: 0, Next: nil}
	expected = &ListNode{Val: 0, Next: nil}
	a.Equal(expected, MergeTwoLists(list1, list2))

	list1 = nil
	list2 = nil
	expected = nil
	a.Equal(expected, MergeTwoLists(list1, list2))
}

func BenchmarkMergeTwoLists(b *testing.B) {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		MergeTwoLists(list1, list2)
	}
}
