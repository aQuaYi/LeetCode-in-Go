package problem1157

// MajorityChecker is ..
type MajorityChecker struct {
	arr []int
}

// Constructor is ...
func Constructor(arr []int) MajorityChecker {
	return MajorityChecker{arr: arr}
}

// Query is ...
func (mc *MajorityChecker) Query(left int, right int, threshold int) int {
	count := make(map[int]int, threshold)
	for i := left; i <= right; i++ {
		count[mc.arr[i]]++
		if count[mc.arr[i]] >= threshold {
			return mc.arr[i]
		}
	}
	return -1
}

/**
 * Your MajorityChecker object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,threshold);
 */
