package problem0887

func superEggDrop(K, N int) int {
	moves := 0
	dp := [101]int{} // 1 <= K <= 100
	// dp[i] = n 表示， i 个鸡蛋，利用 moves 次移动，最多可以检测 n 层楼
	for dp[K] < N {
		for i := K; i > 0; i-- {
			dp[i] += dp[i-1] + 1
			// 以上计算式，是从以下转移方程简化而来
			// dp[moves][k] = 1 + dp[moves-1][k-1] + dp[moves-1][k]
			// 假设 dp[moves-1][k-1] = n0, dp[moves-1][k] = n1
			// 首先检测，从第 n0+1 楼丢下鸡蛋会不会破。
			// 如果鸡蛋破了，F 一定是在 [1：n0] 楼中，
			// 		利用剩下的 moves-1 次机会和 k-1 个鸡蛋，可以把 F 找出来。
			// 如果鸡蛋没破，假如 F 在 [n0+2:n0+n1+1] 楼中
			// 		利用剩下的 moves-1 次机会和 k 个鸡蛋把，也可以把 F 找出来。
			// 所以，当有 moves 个放置机会和 k 个鸡蛋的时候
			// F 在 [1, n0+n1+1] 中的任何一楼，都能够被检测出来。
		}
		moves++
	}
	return moves
}
