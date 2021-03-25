package remove_duplicates_from_sorted_list_2

import "github.com/CNife/leetcode-go/types"

func DeleteDuplicates(head *types.ListNode) *types.ListNode {
	tmpHead := types.ListNode{
		Next: head,
	}
	prev := &tmpHead

	for prev.Next != nil {
		curr, next := prev.Next, prev.Next.Next
		for next != nil && curr.Val == next.Val {
			next = next.Next
		}
		if next == curr.Next {
			prev = prev.Next
		} else {
			prev.Next = next
		}
	}

	return tmpHead.Next
}
