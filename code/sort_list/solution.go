package sort_list

import "github.com/CNife/leetcode-go/types"

func SortList(head *types.ListNode) *types.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	rhs := split(head)
	return merge(SortList(head), SortList(rhs))
}

func split(head *types.ListNode) *types.ListNode {
	ptr := head.Next
	for ptr != nil && ptr.Next != nil {
		head, ptr = head.Next, ptr.Next.Next
	}
	rhs := head.Next
	head.Next = nil
	return rhs
}

var fakeHead = &types.ListNode{Val: -1}

func merge(lhs, rhs *types.ListNode) *types.ListNode {
	ptr := fakeHead
	for lhs != nil && rhs != nil {
		if lhs.Val < rhs.Val {
			ptr.Next, lhs = lhs, lhs.Next
		} else {
			ptr.Next, rhs = rhs, rhs.Next
		}
		ptr = ptr.Next
	}
	if lhs != nil {
		ptr.Next = lhs
	} else {
		ptr.Next = rhs
	}
	return fakeHead.Next
}
