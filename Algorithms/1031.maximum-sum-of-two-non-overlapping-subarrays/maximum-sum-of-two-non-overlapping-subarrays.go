package problem1031

// ref: https://leetcode.com/problems/maximum-sum-of-two-non-overlapping-subarrays/discuss/278251/JavaC%2B%2BPython-O(N)Time-O(1)-Space

func maxSumTwoNoOverlap(A []int, L int, M int) int {
	for i := 1; i < len(A); i++ {
		A[i] += A[i-1]
	}
	// assume original A is A'
	// now, A[i] = sum(A'[:i+1])
	res, lMax, mMax := A[L+M-1], A[L-1], A[M-1]
	for i := L + M; i < len(A); i++ {
		// lMax is max sum of contiguous L elements before the last M elements.
		lMax = max(lMax, A[i-M]-A[i-L-M])
		// mMax is max sum of contiguous M elements before the last L elements.
		mMax = max(mMax, A[i-L]-A[i-L-M])
		res = max(res, max(lMax+A[i]-A[i-M], mMax+A[i]-A[i-L]))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
