package design_circular_deque

type MyCircularDeque struct {
	data       []int
	start, end int
	size       int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		data:  make([]int, k),
		start: 0,
		end:   0,
		size:  0,
	}
}

func (deque *MyCircularDeque) InsertFront(value int) bool {
	if deque.IsFull() {
		return false
	}
	deque.start = deque.leftIndex(deque.start)
	deque.data[deque.start] = value
	deque.size++
	return true
}

func (deque *MyCircularDeque) InsertLast(value int) bool {
	if deque.IsFull() {
		return false
	}
	deque.data[deque.end] = value
	deque.end = deque.rightIndex(deque.end)
	deque.size++
	return true
}

func (deque *MyCircularDeque) DeleteFront() bool {
	if deque.IsEmpty() {
		return false
	}
	deque.start = deque.rightIndex(deque.start)
	deque.size--
	return true
}

func (deque *MyCircularDeque) DeleteLast() bool {
	if deque.IsEmpty() {
		return false
	}
	deque.end = deque.leftIndex(deque.end)
	deque.size--
	return true
}

func (deque *MyCircularDeque) GetFront() int {
	if deque.IsEmpty() {
		return -1
	}
	return deque.data[deque.start]
}

func (deque *MyCircularDeque) GetRear() int {
	if deque.IsEmpty() {
		return -1
	}
	return deque.data[deque.leftIndex(deque.end)]
}

func (deque *MyCircularDeque) IsEmpty() bool {
	return deque.size <= 0
}

func (deque *MyCircularDeque) IsFull() bool {
	return deque.size >= len(deque.data)
}

func (deque *MyCircularDeque) leftIndex(i int) int {
	if i < 1 {
		return len(deque.data) - 1
	}
	return i - 1
}

func (deque *MyCircularDeque) rightIndex(i int) int {
	if i >= len(deque.data)-1 {
		return 0
	}
	return i + 1
}
