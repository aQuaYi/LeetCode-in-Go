package Problem0155

// doublyLinkedNode 是双向链节点
type doublyLinkedNode struct {
	pre, next *doublyLinkedNode
	val       int
}

// MinStack 是可以返回最小值的栈
type MinStack struct {
	min, max *doublyLinkedNode
}

// Constructor 构造 MinStack
func Constructor() MinStack {
	return MinStack{}
}

// Push 存入数据
func (s *MinStack) Push(x int) {
	node := &doublyLinkedNode{val: x}

	switch {
	case s.min == nil:
		s.min, s.max = node, node
	case x <= s.min.val:
		node.next = s.min
		s.min.pre = node
		s.min = node
	case x >= s.max.val:
		s.max.next = node
		node.pre = s.max
		s.max = node
	default:
		temp := s.max
		for temp != nil && temp.val > x {
			temp = temp.pre
		}
		node.next = temp.next
		temp.next.pre = node
		temp.next = node
		node.pre = temp
	}

}

// Pop 抛弃最小值
func (s *MinStack) Pop() {
	if s.min == s.max {
		s.min, s.max = nil, nil
		return
	}
	s.min.next.pre = nil
	s.min = s.min.next
}

// Top 返回最大值
func (s *MinStack) Top() int {
	if s.max != nil {
		return s.max.val
	}

	return 0
}

// GetMin 返回最小值
func (s *MinStack) GetMin() int {
	if s.min != nil {
		return s.min.val
	}

	return 0
}
