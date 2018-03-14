package problem0232

// MyQueue 是利用 栈 实现的队列
type MyQueue struct {
	a, b *Stack
}

// Constructor Initialize your data structure here.
func Constructor() MyQueue {
	return MyQueue{a: NewStack(), b: NewStack()}
}

// Push element x to the back of queue.
func (mq *MyQueue) Push(x int) {
	mq.a.Push(x)
}

// Pop Removes the element from in front of queue and returns that element.
func (mq *MyQueue) Pop() int {
	if mq.b.Len() == 0 {
		for mq.a.Len() > 0 {
			mq.b.Push(mq.a.Pop())
		}
	}

	return mq.b.Pop()
}

// Peek Get the front element.
func (mq *MyQueue) Peek() int {
	res := mq.Pop()
	mq.b.Push(res)
	return res
}

// Empty returns whether the queue is empty.
func (mq *MyQueue) Empty() bool {
	return mq.a.Len() == 0 && mq.b.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

// Stack 是用于存放 int 的 栈
type Stack struct {
	nums []int
}

// NewStack 返回 *kit.Stack
func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

// Push 把 n 放入 栈
func (s *Stack) Push(n int) {
	s.nums = append(s.nums, n)
}

// Pop 从 s 中取出最后放入 栈 的值
func (s *Stack) Pop() int {
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}

// Len 返回 s 的长度
func (s *Stack) Len() int {
	return len(s.nums)
}

// IsEmpty 反馈 s 是否为空
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}
