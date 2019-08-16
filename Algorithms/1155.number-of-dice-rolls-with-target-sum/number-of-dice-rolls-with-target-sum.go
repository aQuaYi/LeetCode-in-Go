package problem1155

const mod = 1e9 + 7

func numRollsToTarget(dices, faces, target int) int {
	dp := [31][1001]int{}
	dp[0][0] = 1

	for d := 1; d <= dices; d++ {
		maxT := min(target, d*faces)
		for t := d; t <= maxT; t++ {
			maxF := min(faces, t)
			for f := 1; f <= maxF; f++ {
				dp[d][t] += dp[d-1][t-f]
			}
			dp[d][t] %= mod
		}
	}

	return dp[dices][target]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
