package problem0918

// ref: https://leetcode.com/problems/maximum-sum-circular-subarray/discuss/178422/C++JavaPython-One-Pass

func maxSubarraySumCircular(A []int) int {
	total, curMax, curMin := 0, 0, 0
	maxSum, minSum := -30001, 30001

	for _, a := range A {
		curMax = max(curMax+a, a)
		maxSum = max(maxSum, curMax)
		curMin = min(curMin+a, a)
		minSum = min(minSum, curMin)
		total += a
	}

	if maxSum > 0 {
		return max(maxSum, total-minSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
