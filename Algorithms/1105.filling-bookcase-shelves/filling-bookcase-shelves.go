package problem1105

func minHeightShelves(books [][]int, width int) int {
	m := len(books)
	dp := make([]int, m+1)

	for i := 1; i <= m; i++ {
		dp[i] = 1000000
		k := 1
		w, h := books[i-k][0], books[i-k][1]
		for w <= width {
			dp[i] = min(dp[i], dp[i-k]+h)
			k++
			if i-k < 0 {
				break
			}
			w += books[i-k][0]
			h = max(h, books[i-k][1])
		}
	}

	return dp[m]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
