// 203. 移除链表元素，简单
// https://leetcode-cn.com/problems/remove-linked-list-elements/

package remove_linked_list_elements

import . "github.com/CNife/leetcode-go/types"

func RemoveElements(head *ListNode, val int) *ListNode {
	dummyHead := ListNode{Next: head}
	prev, curr := &dummyHead, head
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return dummyHead.Next
}
