package problem0225

// MyStack 是用 Queue 实现的 栈
type MyStack struct {
	a, b *Queue
}

// Constructor Initialize your data structure here.
func Constructor() MyStack {
	return MyStack{a: NewQueue(), b: NewQueue()}
}

// Push Push element x onto stack.
func (ms *MyStack) Push(x int) {
	if ms.a.Len() == 0 {
		ms.a, ms.b = ms.b, ms.a
	}
	ms.a.Push(x)
}

// Pop Removes the element on top of the stack and returns that element.
func (ms *MyStack) Pop() int {
	if ms.a.Len() == 0 {
		ms.a, ms.b = ms.b, ms.a
	}

	for ms.a.Len() > 1 {
		ms.b.Push(ms.a.Pop())
	}

	return ms.a.Pop()
}

// Top Get the top element.
func (ms *MyStack) Top() int {
	res := ms.Pop()
	ms.Push(res)
	return res
}

// Empty Returns whether the stack is empty.
func (ms *MyStack) Empty() bool {
	return (ms.a.Len() + ms.b.Len()) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

// Queue 是用于存放 int 的队列
type Queue struct {
	nums []int
}

// NewQueue 返回 *kit.Queue
func NewQueue() *Queue {
	return &Queue{nums: []int{}}
}

// Push 把 n 放入队列
func (q *Queue) Push(n int) {
	q.nums = append(q.nums, n)
}

// Pop 从 q 中取出最先进入队列的值
func (q *Queue) Pop() int {
	res := q.nums[0]
	q.nums = q.nums[1:]
	return res
}

// Len 返回 q 的长度
func (q *Queue) Len() int {
	return len(q.nums)
}

// IsEmpty 反馈 q 是否为空
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
