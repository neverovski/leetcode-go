package problem141

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	a := assert.New(t)

	l1 := &ListNode{Val: 3}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 0}
	l4 := &ListNode{Val: -4}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l2

	a.Equal(true, hasCycle(l1))
}
