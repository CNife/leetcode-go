package my_calendar_1

import "sort"

type MyCalendar struct {
	nodes []intervalNode
}

type intervalNode struct {
	time    int
	isStart bool
}

func Constructor() MyCalendar {
	return MyCalendar{nil}
}

func (c *MyCalendar) Clear() {
	c.nodes = nil
}

func (c *MyCalendar) Book(start, end int) bool {
	i := sort.Search(len(c.nodes), func(i int) bool {
		return start < c.nodes[i].time
	})
	if i >= len(c.nodes) {
		c.nodes = append(c.nodes, intervalNode{start, true}, intervalNode{end, false})
		return true
	}
	if c.nodes[i].isStart && end <= c.nodes[i].time {
		c.nodes = append(c.nodes[:i+2], c.nodes[i:]...)
		c.nodes[i] = intervalNode{start, true}
		c.nodes[i+1] = intervalNode{end, false}
		return true
	}
	return false
}
