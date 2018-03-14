package problem0155

// MinStack 是可以返回最小值的栈
type MinStack struct {
	stack []item
}
type item struct {
	min, x int
}

// Constructor 构造 MinStack
func Constructor() MinStack {
	return MinStack{}
}

// Push 存入数据
func (s *MinStack) Push(x int) {
	min := x
	if len(s.stack) > 0 && s.GetMin() < x {
		min = s.GetMin()
	}
	s.stack = append(s.stack, item{min: min, x: x})
}

// Pop 抛弃最后一个入栈的值
func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
}

// Top 返回最大值
func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1].x
}

// GetMin 返回最小值
func (s *MinStack) GetMin() int {
	return s.stack[len(s.stack)-1].min
}
