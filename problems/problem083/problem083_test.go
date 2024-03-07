package problem083

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicatesUsingOneList(t *testing.T) {
	a := assert.New(t)

	a.Nil(deleteDuplicatesUsingOneList(nil))
	a.Equal(&ListNode{}, deleteDuplicatesUsingOneList(&ListNode{}))
	a.Equal(&ListNode{1, &ListNode{2, nil}}, deleteDuplicatesUsingOneList(&ListNode{1, &ListNode{1, &ListNode{2, nil}}}))
	a.Equal(&ListNode{1, &ListNode{2, &ListNode{3, nil}}}, deleteDuplicatesUsingOneList(&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}))
}

func TestDeleteDuplicatesUsingTwoList(t *testing.T) {
	a := assert.New(t)

	a.Nil(deleteDuplicatesUsingTwoList(nil))
	a.Equal(&ListNode{}, deleteDuplicatesUsingTwoList(&ListNode{}))
	a.Equal(&ListNode{1, &ListNode{2, nil}}, deleteDuplicatesUsingTwoList(&ListNode{1, &ListNode{1, &ListNode{2, nil}}}))
	a.Equal(&ListNode{1, &ListNode{2, &ListNode{3, nil}}}, deleteDuplicatesUsingTwoList(&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}))
}

func BenchmarkDeleteDuplicatesUsingOneList(b *testing.B) {
	head := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		deleteDuplicatesUsingOneList(head)
	}
}

func BenchmarkDeleteDuplicatesUsingTwoList(b *testing.B) {
	head := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		deleteDuplicatesUsingTwoList(head)
	}
}
