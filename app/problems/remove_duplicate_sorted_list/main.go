package removeduplicatesortedlist

// Problem: https://leetcode.com/problems/remove-duplicates-from-sorted-list/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var curr = head
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
