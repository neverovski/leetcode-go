package problem160

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIntersectionNode(t *testing.T) {
	a := assert.New(t)

	// Create intersecting linked lists
	intersectNode := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}
	headA := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: intersectNode}}
	headB := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: intersectNode}}}

	// Test intersecting linked lists
	result := getIntersectionNode(headA, headB)
	a.Equal(intersectNode, result, "Should return the intersecting node")

	// Create non-intersecting linked lists
	headA = &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: nil}}}
	headB = &ListNode{Val: 3, Next: nil}

	// Test non-intersecting linked lists
	result = getIntersectionNode(headA, headB)
	a.Nil(result, "Should return nil for non-intersecting linked lists")
}
