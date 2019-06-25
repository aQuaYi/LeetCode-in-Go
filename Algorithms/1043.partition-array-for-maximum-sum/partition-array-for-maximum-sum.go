package problem1043

func maxSumAfterPartitioning(A []int, K int) int {
	i := indexOfMax(A)
	n := len(A)
	if n <= K {
		return n * A[i]
	}
	L, R := max(0, i-K+1), min(n-1, i+K-1) // [L,R]
	mL, mR, mSum := L, R, 0
	for l := L; l <= i && l+K-1 <= R; l++ {
		r := l + K - 1
		sum := 0
		for j := L; j <= R; j++ {
			if l <= j && j <= r {
				continue
			}
			sum += A[j]
		}
		if mSum < sum {
			mSum, mL, mR = sum, l, r
		}
	}
	left := 0
	if mL > 0 {
		left = maxSumAfterPartitioning(A[:mL], K)
	}

	right := 0
	if mR < n-1 {
		right = maxSumAfterPartitioning(A[mR+1:], K)
	}
	return A[i]*K + left + right
}

func indexOfMax(A []int) int {
	a, i := A[0], 0
	for j := 1; j < len(A); j++ {
		if a < A[j] {
			a, i = A[j], j
		}
	}
	return i
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
