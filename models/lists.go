package models

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = MergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = MergeTwoLists(list1, list2.Next)
	return list2
}

func ConvertToLinkedList(arr []int) *ListNode {
	var head, current *ListNode
	for _, val := range arr {
		node := &ListNode{Val: val}
		if head == nil {
			head = node
			current = node
		} else {
			current.Next = node
			current = node
		}
	}
	return head
}

func ConvertLinkedListToArray(head *ListNode) []int {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return arr
}
