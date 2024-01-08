package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	// Test case: MergeTwoLists - Successful case
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}}
	list2 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}

	expectedMergedList := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}

	merged := MergeTwoLists(list1, list2)

	assert.Equal(t, expectedMergedList, merged)
}
