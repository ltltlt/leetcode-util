package listutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(elems []int) *ListNode {
	if len(elems) == 0 {
		return nil
	}

	head := &ListNode{Val: elems[0]}
	current := head
	for _, elem := range elems[1:] {
		current.Next = &ListNode{Val: elem}
		current = current.Next
	}

	return head
}

func ListMatchSlice(t *testing.T, head *ListNode, elems []int) {
	current := head
	for i, elem := range elems {
		assert.NotNilf(t, current, "element %d is nil", i)
		assert.Equalf(t, elem, current.Val, "element %d not match", i)
		current = current.Next
	}
	assert.Nil(t, current)
}

func ListEqual(t *testing.T, head1, head2 *ListNode) {
	ptr1, ptr2 := head1, head2
	i := 0
	for ptr1 != nil && ptr2 != nil {
		assert.Equalf(t, ptr1.Val, ptr2.Val, "element %d not match", i)
		ptr1, ptr2 = ptr1.Next, ptr2.Next
		i++
	}
	assert.Equal(t, ptr1, ptr2)
}
