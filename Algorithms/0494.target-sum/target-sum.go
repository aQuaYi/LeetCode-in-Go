package problem0494

func findTargetSumWays(nums []int, S int) int {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum < S {
		return 0
	}
	// sum 和 S 必须同为奇数或者偶数
	if (sum+S)%2 == 1 {
		return 0
	}

	return calcSumWays(nums, (sum+S)/2)
}

// nums 被分为两个部分 P 和 N
// P 中的元素前面放的是 +
// N 中的元素前面放的是 -
// 可得
//                   sum(P) - sum(N) = target
// sum(P) + sum(N) + sum(P) - sum(N) = target + sum(P) + sum(N)
//                        2 * sum(P) = target + sum(nums)
// 于是，题目就被转换成了
// nums 有多少个子集，其和为 (target + sum(nums))/2
func calcSumWays(nums []int, target int) int {
	// dp[i] 表示 nums 中和为 i 的子集个数
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[target]
}
