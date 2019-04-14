package problem0239

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	window := make([]int, 0, k)
	for i, v := range nums {
		// 控制 window 窗口大小
		if i > k-1 && window[0] == i-k {
			window = window[1:]
		}

		// window 保持递减
		j := len(window) - 1
		for ; j >= 0; j-- {
			if nums[window[j]] > v {
				break
			}
		}
		window = window[:j+1]
		window = append(window, i)

		// 获取 window 最大值
		if i >= k-1 {
			res = append(res, nums[window[0]])
		}
	}

	return res
}

// 解法二
// 参看 https://leetcode.com/problems/sliding-window-maximum/discuss/65881/O(n)-solution-in-Java-with-two-simple-pass-in-the-array
// 例子
// nums 1 3 3 -3 5 5 6 7
// k = 3
// left  1 3 3 | -3 5 5 | 6 7
// right 3 3 -1|  5 5 3 | 7 7
// res   3 3 5    5 6 7
func maxSlidingWindow2(nums []int, k int) []int {
	size := len(nums)
	if size <= 1 {
		return nums
	}

	left, right := make([]int, size), make([]int, size)
	left[0], right[size-1] = nums[0], nums[size-1]

	for i := 1; i < size; i++ {
		if i%k == 0 {
			left[i] = nums[i]
		} else {
			left[i] = max(nums[i], left[i-1])
		}

		j := size - 1 - i
		if (j+1)%k == 0 {
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
