package Problem0155

// MinStack 是可以返回最小值的栈
type MinStack struct {
	min   int
	stack []int
}

// Constructor 构造 MinStack
func Constructor() MinStack {
	return MinStack{}
}

// Push 存入数据
func (s *MinStack) Push(x int) {
	if len(s.stack) == 0 {
		s.min = x
	}

	s.stack = append(s.stack, x)
	if s.min > x {
		s.min = x
	}
}

// Pop 抛弃最后一个入栈的值
func (s *MinStack) Pop() {
	if len(s.stack) == 0 {
		return
	}

	s.stack = s.stack[:len(s.stack)-1]
	s.min = 1<<63 - 1
	for i := range s.stack {
		if s.min > s.stack[i] {
			s.min = s.stack[i]
		}
	}
}

// Top 返回最大值
func (s *MinStack) Top() int {
	if len(s.stack) > 0 {
		return s.stack[len(s.stack)-1]
	}
	return 0
}

// GetMin 返回最小值
func (s *MinStack) GetMin() int {
	return s.min
}
