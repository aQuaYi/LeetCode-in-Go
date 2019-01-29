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
func (this *MinStack) Push(x int) {
	min := x
	if len(this.stack) > 0 && this.GetMin() < x {
		min = this.GetMin()
	}
	this.stack = append(this.stack, item{min: min, x: x})
}

// Pop 抛弃最后一个入栈的值
func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

// Top 返回最大值
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].x
}

// GetMin 返回最小值
func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}
