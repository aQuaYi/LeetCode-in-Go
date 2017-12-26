package Problem0715

// RangeModule 记录了跟踪的范围
type RangeModule struct {
}

// Constructor 返回新建的 RangeModule
func Constructor() RangeModule {

	return RangeModule{}
}

// AddRange 添加追踪的返回
func (r *RangeModule) AddRange(left int, right int) {

}

// QueryRange 返回 true 如果 [left, right) 全部都在追踪范围内
func (r *RangeModule) QueryRange(left int, right int) bool {

	return true
}

// RemoveRange 从追踪范围中删除 [left,right)
func (r *RangeModule) RemoveRange(left int, right int) {

}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
