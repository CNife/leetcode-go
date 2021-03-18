package reverse_linked_list_2

import "github.com/CNife/leetcode-go/types"

func ReverseBetween(head *types.ListNode, left, right int) *types.ListNode {
	if head == nil || left <= 0 || right <= left {
		return head
	}

	var prev *types.ListNode
	if left == 1 {
		prev = &types.ListNode{Next: head}
	} else {
		prev = head
		for i := 2; i < left; i++ {
			prev = prev.Next
		}
	}

	node := prev.Next
	for i := left; i < right; i++ {
		next := node.Next
		node.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	if left == 1 {
		return prev.Next
	} else {
		return head
	}
}
