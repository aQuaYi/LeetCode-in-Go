package problem0813

var avgs = [101][101]float64{}

func largestSumOfAverages(A []int, K int) float64 {
	if K == 1 {
		return avg(A)
	}

	n := len(A)

	sums := make([]int, n+1)
	for i := range A {
		sums[i+1] = sums[i] + A[i]
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			avgs[i][j] = float64(sums[j]-sums[i]) / float64(j-i)
		}
	}

	return helper(n, 0, K)
}

func helper(n, i, k int) float64 {
	if k == 1 {
		return avgs[i][n]
	}

	maxAvg := -1.
	for j := i + 1; j <= n-k+1; j++ {
		maxAvg = max(maxAvg, avgs[i][j]+helper(n, j, k-1))
	}

	return maxAvg
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func avg(A []int) float64 {
	sum := 0
	for i := range A {
		sum += A[i]
	}
	return float64(sum) / float64(len(A))
}
