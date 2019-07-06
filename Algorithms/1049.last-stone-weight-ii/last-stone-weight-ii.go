package problem1049

func lastStoneWeightII(stones []int) int {
	dp := [3001]bool{}
	dp[0] = true
	sum := 0
	for _, s := range stones {
		sum += s
		for i := sum; i >= s; i-- {
			// 把所有可能的 sum 组合都记录下来
			dp[i] = dp[i] || dp[i-s]
		}
	}
	// 此时 sum 是 sum(stones)
	// stones 分成两组后，重量分别是 part 和 sum-part
	// 答案是 sum-part - part
	// 最好的情况是 part = sum/2
	// 如果不行的话，就看看 part-- 行不行
	// 以此类推直到 part = 0
	part := sum / 2
	for ; part >= 0; part-- {
		if dp[part] {
			break
		}
	}
	return sum - part - part
}
