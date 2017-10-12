package Problem0295

// MedianFinder 用于寻找 Median
type MedianFinder struct {
	isEven      bool
	left, right *heap
}

// Constructor initialize your data structure here.
func Constructor() MedianFinder {
	isEven := true
	left := newHeap(less)
	right := newHeap(bigger)

	return MedianFinder{
		isEven: isEven,
		left:   left,
		right:  right,
	}
}

// AddNum 给 MedianFinder 添加数
func (mf *MedianFinder) AddNum(n int) {
	if mf.isEven && n < mf.left.Root() {
		mf.left.Add(n)
	}

	mf.isEven = !mf.isEven
}

// FindMedian 给出 Median
func (mf *MedianFinder) FindMedian() float64 {
	if mf.isEven {
		return float64(mf.left.Root()+mf.right.Root()) / 2
	}
	return float64(mf.left.Root())
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type heap struct {
	nums []int
	less func(int, int) bool
}

// newHeap 返回 *heap
func newHeap(compareFunc func(int, int) bool) *heap {
	// nums[0] 保留不用
	nums := make([]int, 1, 1024)
	return &heap{
		nums: nums,
		less: compareFunc,
	}
}

func (h *heap) Add(n int) {
	k := len(h.nums)
	h.nums = append(h.nums, n)
	for k > 1 && h.less(h.nums[k/2], h.nums[k]) {
		h.nums[k/2], h.nums[k] = h.nums[k], h.nums[k/2]
		k /= 2
	}
}

func (h *heap) Root() int {
	return h.nums[1]
}

func less(a, b int) bool {
	return a < b
}

func bigger(a, b int) bool {
	return a > b
}
