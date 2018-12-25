package problem0930

// ref: https://leetcode.com/problems/binary-subarrays-with-sum/discuss/186683/C%2B%2BJavaPython-Straight-Forward

func numSubarraysWithSum(A []int, S int) int {
	preSum, res := 0, 0
	count := make([]int, len(A)+1)
	count[0] = 1
	for _, n := range A {
		preSum += n
		if preSum >= S {
			res += count[preSum-S]
		}
		count[preSum]++
	}
	return res
}
