package implement_queue_using_stacks

type MyQueue struct {
	in, out stack
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.in.push(x)
}

func (q *MyQueue) Pop() int {
	q.prepareForOut()

	if len(q.out) < 1 {
		return -1
	}
	return q.out.pop()
}

func (q *MyQueue) Peek() int {
	q.prepareForOut()

	if len(q.out) < 1 {
		return -1
	}
	return q.out.peek()
}

func (q *MyQueue) Empty() bool {
	return len(q.in) < 1 && len(q.out) < 1
}

func (q *MyQueue) prepareForOut() {
	if len(q.out) < 1 {
		for len(q.in) > 0 {
			q.out.push(q.in.pop())
		}
	}
}

type stack []int

func (s *stack) push(x int) {
	*s = append(*s, x)
}

func (s *stack) pop() int {
	n := len(*s) - 1
	x := (*s)[n]
	*s = (*s)[:n]
	return x
}

func (s stack) peek() int {
	return s[len(s)-1]
}
