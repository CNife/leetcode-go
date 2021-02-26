package reorder_list

import "github.com/CNife/leetcode/go/types"

func ReorderList(head *types.ListNode) {
	if head != nil {
		insert(head, reverse(split(head)))
	}
}

func split(head *types.ListNode) *types.ListNode {
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
			slow = slow.Next
		}
	}
	rightHead := slow.Next
	slow.Next = nil
	return rightHead
}

func reverse(head *types.ListNode) *types.ListNode {
	var prev *types.ListNode = nil
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func insert(lhs, rhs *types.ListNode) {
	for rhs != nil {
		lhsNext, rhsNext := lhs.Next, rhs.Next
		lhs.Next, rhs.Next = rhs, lhsNext
		lhs, rhs = lhsNext, rhsNext
	}
}
