package problem1186

func maximumSum(A []int) int {
	sum, curMin := 0, 1000000
	res, maxA := 0, -1000000
	for _, a := range A {
		sum += a
		curMin = min(curMin, a)

		if sum-min(curMin, 0) <= 0 {
			sum = 0
			curMin = 1000000
		} else {
			res = max(res, sum-min(curMin, 0))
		}

		maxA = max(maxA, a)
	}

	if maxA <= 0 {
		return maxA
	}

return max(res, sum-min(curMin,0))
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
