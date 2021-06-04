// 160. 相交链表，简单
// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/

package intersection_of_two_linked_lists

import . "github.com/CNife/leetcode-go/types"

func GetIntersectionNode(lhs, rhs *ListNode) *ListNode {
	ll, rl := goToLast(lhs), goToLast(rhs)
	if ll.node != rl.node {
		return nil
	}
	if lll, rll := ll.length, rl.length; lll > rll {
		for i := rll; i < lll; i++ {
			lhs = lhs.Next
		}
	} else if rll > lll {
		for i := lll; i < rll; i++ {
			rhs = rhs.Next
		}
	}
	for lhs != rhs {
		lhs, rhs = lhs.Next, rhs.Next
	}
	return lhs
}

type listLast struct {
	node   *ListNode
	length int
}

func goToLast(list *ListNode) listLast {
	node, length := list, 0
	for node != nil {
		node = node.Next
		length++
	}
	return listLast{node, length}
}
