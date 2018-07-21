package problem0622

// MyCircularQueue 结构体
type MyCircularQueue struct {
	queue []int
	k     int
}

// Constructor initialize your data structure here. Set the size of the queue to be k.
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, 0, k*3),
		k:     k,
	}
}

// EnQueue insert an element into the circular queue. Return true if the operation is successful.
func (m *MyCircularQueue) EnQueue(value int) bool {
	if len(m.queue) == m.k {
		return false
	}
	m.queue = append(m.queue, value)
	return true
}

// DeQueue delete an element from the circular queue. Return true if the operation is successful.
func (m *MyCircularQueue) DeQueue() bool {
	if len(m.queue) == 0 {
		return false
	}

	m.queue = m.queue[1:]
	return true
}

// Front get the front item from the queue.
func (m *MyCircularQueue) Front() int {
	if len(m.queue) == 0 {
		return -1
	}

	return m.queue[0]
}

// Rear get the last item from the queue. */
func (m *MyCircularQueue) Rear() int {
	if len(m.queue) == 0 {
		return -1
	}
	return m.queue[len(m.queue)-1]
}

// IsEmpty checks whether the circular queue is empty or not. */
func (m *MyCircularQueue) IsEmpty() bool {
	return len(m.queue) == 0
}

// IsFull checks whether the circular queue is full or not. */
func (m *MyCircularQueue) IsFull() bool {
	return len(m.queue) == m.k
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
