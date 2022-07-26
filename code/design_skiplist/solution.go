package design_skiplist

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type SkipList struct {
	head *Node
}

type Node struct {
	val         int
	right, down *Node
}

func Constructor() SkipList {
	return SkipList{
		head: &Node{
			val: -1,
		},
	}
}

func (sl *SkipList) Search(target int) bool {
	for ptr := sl.head; ptr != nil; {
		for ptr.right != nil && ptr.right.val < target {
			ptr = ptr.right
		}
		if ptr.right == nil || ptr.right.val > target {
			ptr = ptr.down
		} else {
			return true
		}
	}
	return false
}

func (sl *SkipList) Add(num int) {
	var path []*Node
	for ptr := sl.head; ptr != nil; {
		for ptr.right != nil && ptr.right.val < num {
			ptr = ptr.right
		}
		path = append(path, ptr)
		ptr = ptr.down
	}

	insertUp := true
	var down *Node = nil
	for insertUp && len(path) > 0 {
		insert := path[len(path)-1]
		path = path[:len(path)-1]
		insert.right = &Node{
			val:   num,
			right: insert.right,
			down:  down,
		}
		down = insert.right
		insertUp = rand.Intn(2) == 1
	}

	if insertUp {
		sl.head = &Node{
			val: -1,
			right: &Node{
				val:  num,
				down: down,
			},
			down: sl.head,
		}
	}
}

func (sl *SkipList) Erase(num int) bool {
	seen := false
	for ptr := sl.head; ptr != nil; {
		for ptr.right != nil && ptr.right.val < num {
			ptr = ptr.right
		}
		if ptr.right == nil || ptr.right.val > num {
			ptr = ptr.down
		} else {
			seen = true
			ptr.right = ptr.right.right
			ptr = ptr.down
		}
	}
	return seen
}
