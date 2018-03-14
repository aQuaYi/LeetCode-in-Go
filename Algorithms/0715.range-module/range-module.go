package problem0715

import (
	"sort"
)

// RangeModule 记录了跟踪的范围
type RangeModule struct {
	ranges []*interval
}

// Constructor 返回新建的 RangeModule
func Constructor() RangeModule {
	return RangeModule{ranges: make([]*interval, 0, 2048)}
}

// AddRange 添加追踪的返回
func (r *RangeModule) AddRange(left int, right int) {
	it := &interval{left: left, right: right}

	n := len(r.ranges)
	i := sort.Search(n, func(i int) bool {
		return left <= r.ranges[i].right
	})

	var j int
	for j = i; j < n && r.ranges[j].left <= right; j++ {
		it.add(r.ranges[j])
	}

	if i == j {
		r.ranges = append(r.ranges, nil)
	}

	copy(r.ranges[i+1:], r.ranges[j:])
	r.ranges = r.ranges[:n-j+i+1]
	r.ranges[i] = it
}

// QueryRange 返回 true 如果 [left, right) 全部都在追踪范围内
func (r *RangeModule) QueryRange(left int, right int) bool {
	n := len(r.ranges)
	i := sort.Search(n, func(i int) bool {
		return right <= r.ranges[i].right
	})
	return 0 <= i && i < n && r.ranges[i].isCover(left, right)
}

// RemoveRange 从追踪范围中删除 [left,right)
func (r *RangeModule) RemoveRange(left int, right int) {
	it := &interval{left: left, right: right}

	n := len(r.ranges)
	if n == 0 {
		return
	}

	i := sort.Search(n, func(i int) bool {
		return left < r.ranges[i].right
	})

	temp := make([]*interval, 0, 16)
	var j int
	for j = i; j < n && r.ranges[j].left < right; j++ {
		ra, rb := minus(r.ranges[j], it)
		if ra != nil {
			temp = append(temp, ra)
		}
		if rb != nil {
			temp = append(temp, rb)
		}
	}

	if i == j {
		return
	}

	r.ranges = append(r.ranges, nil)
	copy(r.ranges[i+len(temp):], r.ranges[j:])
	r.ranges = r.ranges[:n-j+i+len(temp)]

	for k := 0; k < len(temp); k++ {
		r.ranges[i+k] = temp[k]
	}
}

type interval struct {
	left, right int
}

func (it *interval) isCover(left, right int) bool {
	return it.left <= left && right <= it.right
}

func (it *interval) add(a *interval) {
	if a.left < it.left {
		it.left = a.left
	}

	if it.right < a.right {
		it.right = a.right
	}
}

// 返回 a-b
func minus(a, b *interval) (*interval, *interval) {
	if b.left <= a.left && a.right <= b.right {
		return nil, nil
	}

	if b.left <= a.left && a.left < b.right && b.right < a.right {
		return &interval{left: b.right, right: a.right}, nil
	}

	if a.left < b.left && b.left < a.right && a.right < b.right {
		return &interval{left: a.left, right: b.left}, nil
	}

	// a.left < b.left && b.right < a.right
	return &interval{left: a.left, right: b.left},
		&interval{left: b.right, right: a.right}

}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
