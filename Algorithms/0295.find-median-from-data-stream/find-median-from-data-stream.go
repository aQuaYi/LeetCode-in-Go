package Problem0295

import (
	"sort"
)

// MedianFinder 用于寻找 Median
type MedianFinder struct {
	nums []int
}

// Constructor initialize your data structure here.
func Constructor() MedianFinder {
	return MedianFinder{[]int{}}
}

// AddNum 给 MedianFinder 添加数
func (mf *MedianFinder) AddNum(n int) {
	mf.nums = append(mf.nums, n)
}

// FindMedian 给出 Median
func (mf *MedianFinder) FindMedian() float64 {
	sort.Ints(mf.nums)
	i, j := (len(mf.nums)-1)/2, len(mf.nums)/2
	return float64(mf.nums[i]+mf.nums[j]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
