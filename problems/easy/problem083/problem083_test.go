package problem083

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  *ListNode
	out *ListNode
}

func TestDeleteDuplicatesUsingOneList(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in:  nil,
			out: nil,
		},
		{
			in:  &ListNode{},
			out: &ListNode{},
		},
		{
			in:  &ListNode{1, &ListNode{1, &ListNode{2, nil}}},
			out: &ListNode{1, &ListNode{2, nil}},
		},
		{
			in:  &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}},
			out: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		},
	}

	for key, tc := range testCases {
		if tc.in == nil {
			a.Nil(deleteDuplicatesUsingOneList(tc.out), fmt.Sprintf("TestDeleteDuplicatesUsingOneList number # %d", key+1))
		} else {
			a.Equal(tc.out, deleteDuplicatesUsingOneList(tc.in), fmt.Sprintf("TestDeleteDuplicatesUsingOneList number # %d", key+1))
		}
	}
}

func TestDeleteDuplicatesUsingTwoList(t *testing.T) {
	a := assert.New(t)

	testCases := []TestCase{
		{
			in:  nil,
			out: nil,
		},
		{
			in:  &ListNode{},
			out: &ListNode{},
		},
		{
			in:  &ListNode{1, &ListNode{1, &ListNode{2, nil}}},
			out: &ListNode{1, &ListNode{2, nil}},
		},
		{
			in:  &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}},
			out: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		},
	}

	for key, tc := range testCases {
		if tc.in == nil {
			a.Nil(deleteDuplicatesUsingTwoList(tc.out), fmt.Sprintf("TestDeleteDuplicatesUsingTwoList number # %d", key+1))
		} else {
			a.Equal(tc.out, deleteDuplicatesUsingTwoList(tc.in), fmt.Sprintf("TestDeleteDuplicatesUsingTwoList number # %d", key+1))
		}
	}
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
