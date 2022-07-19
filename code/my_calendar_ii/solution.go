package my_calendar_2

type MyCalendarTwo struct {
	root segTreeNode
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		root: segTreeNode{},
	}
}

func (c *MyCalendarTwo) Book(start, end int) bool {
	if c.root.query(0, N, start, end-1) == 2 {
		return false
	}
	c.root.update(0, N, start, end-1, 1)
	return true
}

const N int = 1e9

type segTreeNode struct {
	left, right *segTreeNode
	value, add  int
}

func (n *segTreeNode) update(start, end, l, r, value int) {
	if l <= start && end <= r {
		n.value += value
		n.add += value
		return
	}

	n.pushDown()
	middle := start + (end-start)/2
	if l <= middle {
		n.left.update(start, middle, l, r, value)
	}
	if r > middle {
		n.right.update(middle+1, end, l, r, value)
	}
	n.pushUp()
}

func (n *segTreeNode) query(start, end, l, r int) int {
	if l <= start && end <= r {
		return n.value
	}

	n.pushDown()
	middle := start + (end-start)/2
	result := 0
	if l <= middle {
		result = n.left.query(start, middle, l, r)
	}
	if r > middle {
		result = max(result, n.right.query(middle+1, end, l, r))
	}
	return result
}

func (n *segTreeNode) pushDown() {
	if n.left == nil {
		n.left = &segTreeNode{}
	}
	if n.right == nil {
		n.right = &segTreeNode{}
	}
	if n.add == 0 {
		return
	}
	n.left.value += n.add
	n.right.value += n.add
	n.left.add += n.add
	n.right.add += n.add
	n.add = 0
}

func (n *segTreeNode) pushUp() {
	n.value = max(n.left.value, n.right.value)
}

func max(lhs, rhs int) int {
	if lhs < rhs {
		return rhs
	}
	return lhs
}
