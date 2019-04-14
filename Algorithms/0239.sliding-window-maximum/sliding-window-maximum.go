package problem0239

// 参看 https://leetcode.com/problems/sliding-window-maximum/discuss/65881/O(n)-solution-in-Java-with-two-simple-pass-in-the-array
// 例子
// nums = [2,1,3,4,6,3,8,9,10,12,56], k=4
// split   2, 1, 3, 4 | 6, 3,   8,  9 | 10, 12, 56  // every k number into a group
// left    2, 2, 3, 4 | 6, 6,   8,  9 | 10, 12, 56
// right   4, 4, 4, 4 | 9, 9,   9,  9 | 56, 56, 56
// res[i] = max(right_max[i], left[i+w-1])
// res     4, 6, 6, 8,| 9, 10, 12, 56
func maxSlidingWindow(nums []int, k int) []int {
	size := len(nums)
	if k <= 1 {
		return nums
	}

	g := k - 1

	left := make([]int, size)
	left[0] = nums[0]
	for i := 1; i < size; i++ {
		if i%g == 1 {
			left[i] = nums[i]
		} else {
			left[i] = max(nums[i], left[i-1])
		}
	}

	right := make([]int, size)
	right[size-1] = nums[size-1]
	for j := size - 2; j >= 0; j-- {
		if j%g == 0 {
			right[j] = nums[j]
		} else {
			right[j] = max(nums[j], right[j+1])
		}
	}

	res := make([]int, size-k+1)
	for i := 0; i <= size-k; i++ {
		res[i] = max(right[i], left[i+k-1])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
