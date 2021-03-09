package remove_nth_node_from_end_of_list

import "github.com/CNife/leetcode-go/types"

func RemoveNthFromEnd(head *types.ListNode, n int) *types.ListNode {
	if n < 1 {
		return head
	}

	ahead := head
	for i := 0; i < n && ahead != nil; i++ {
		ahead = ahead.Next
	}
	if ahead == nil {
		return head.Next
	}

	behind := head
	for ahead.Next != nil {
		ahead = ahead.Next
		behind = behind.Next
	}

	behind.Next = behind.Next.Next
	return head
}
