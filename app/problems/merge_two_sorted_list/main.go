package mergetwosortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// Problem: https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var head *ListNode
	var curr *ListNode

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		head = list1
		curr = list1
		list1 = list1.Next
	} else {
		head = list2
		curr = list2
		list2 = list2.Next
	}

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			curr = curr.Next
			list1 = list1.Next
		} else {
			curr.Next = list2
			curr = curr.Next
			list2 = list2.Next
		}
	}
	for list1 != nil {
		curr.Next = list1
		curr = curr.Next
		list1 = list1.Next
	}

	for list2 != nil {
		curr.Next = list2
		curr = curr.Next
		list2 = list2.Next
	}

	return head
}
