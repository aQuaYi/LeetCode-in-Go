package problem0805

// n = len(A)
// k = len(B)
// sum = A 的所有数之和
// s   = B 的所有数之和
// 可由 s/k==(sum-s)/(n-k)
// 得 sum/n==s/k
// 即 s*n == sum*k
// 这是返回 true 的充分必要条件
//
// 根据对称性，可以假设 len(B)<=len(C)
// 可得 k<=n-k,
// k 的范围是 1<=k<=n/2
func splitArraySameAverage(A []int) bool {
	n := len(A)
	sum := 0
	for _, num := range A {
		sum += num
	}

	dp := make([][]bool, n/2+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}

	dp[0][0] = true
	// dp[k][s] == true 的含义是
	// 当 B 是由 k 个元素组成时，s 是其中一个可能的 sum 值
	// 所以 dp[0][0] == true 表示
	// B 由 0 个元素组成时，其和为 0
	for _, num := range A {
		for s := sum; s >= num; s-- {
			for k := 1; k <= n/2; k++ {
				// 当把 num 作为第 k 个数，添加到 B 以后，
				// B 的总和为 s
				// 此时，dp[s][k] 想要为 true
				// 要么，由别的 k 个数之和为 s     的组合成的 B 为 true
				// 要么，由  k-1 个数之和为 s-num  的组合成的 B 为 true
				dp[k][s] = dp[k][s] || dp[k-1][s-num]

				// 如果 k 和 s 是一个可行的组合，
				// 就检查 sum/n =?= s/k
				// 两者都成立，说明找到了题目要求的组合
				// 立即返回 true
				if dp[k][s] && sum*k == s*n {
					return true
				}
			}
		}
	}

	return false
}
