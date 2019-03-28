package problem0232

import "container/list"

// MyQueue 是利用 list 实现的队列
type MyQueue struct {
	list *list.List
}

// Constructor Initialize your data structure here.
func Constructor() MyQueue {
	return MyQueue{
		list: list.New(),
	}
}

// Push element x to the back of queue.
func (q *MyQueue) Push(x int) {
	q.list.PushBack(x)
}

// Pop Removes the element from in front of queue and returns that element.
func (q *MyQueue) Pop() int {
	front := q.list.Front()
	res := front.Value.(int)
	q.list.Remove(front)
	return res
}

// Peek Get the front element.
func (q *MyQueue) Peek() int {
	front := q.list.Front()
	res := front.Value.(int)
	return res
}

// Empty returns whether the queue is empty.
func (q *MyQueue) Empty() bool {
	return q.list.Len() == 0
}
