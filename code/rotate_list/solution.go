package rotate_list

import "github.com/CNife/leetcode-go/types"

func RotateRight(head *types.ListNode, k int) *types.ListNode {
	if head == nil || k <= 0 {
		return head
	}

	last, size := traverseList(head)

	index := size - 1 - k%size
	if index == size-1 {
		return head
	}

	newTail := getListNode(head, index)
	newHead := newTail.Next
	newTail.Next, last.Next = nil, head
	return newHead
}

func traverseList(head *types.ListNode) (last *types.ListNode, size int) {
	for last, size = head, 1; last.Next != nil; last = last.Next {
		size++
	}
	return
}

func getListNode(head *types.ListNode, index int) *types.ListNode {
	for i := 0; i < index; i++ {
		head = head.Next
	}
	return head
}
