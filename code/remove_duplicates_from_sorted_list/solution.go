package remove_duplicates_from_sorted_list

import "github.com/CNife/leetcode-go/types"

func DeleteDuplicates(head *types.ListNode) *types.ListNode {
	if head == nil {
		return nil
	}

	c, p := head, head.Next
	for p != nil {
		if p.Val != c.Val {
			if p != c.Next {
				c.Next = p
			}
			c = p
		}
		p = p.Next
	}

	c.Next = nil
	return head
}
