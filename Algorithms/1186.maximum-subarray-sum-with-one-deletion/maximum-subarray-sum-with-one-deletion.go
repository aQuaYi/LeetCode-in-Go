package problem1186

func maximumSum(A []int) int {
	res, maxA := 0, -1000000
	last, cur := 0, 0
	for _, a := range A {
		if a < 0 {
			res = max(res, last+cur)
			last, cur = cur, 0
		}else{
			cur += a
		}
		maxA= max(maxA, a)
	}

	if maxA<=0 {
			return maxA
	}

	return max(res, last+cur)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
