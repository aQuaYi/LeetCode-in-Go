package problem0813

func largestSumOfAverages(A []int, K int) float64 {
	if K == 1 {
		return avg(A)
	}
	sum := 0.0
	maxAvg := -1.
	for i := 0; i < len(A)-K+1; i++ {
		sum += float64(A[i])
		maxAvg = max(maxAvg, sum/float64(i+1)+largestSumOfAverages(A[i+1:], K-1))
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
