package problem0962

func maxWidthRamp(A []int) int {
	s := []int{}
	for i := range A {
		if len(s) == 0 || A[i] < A[s[len(s)-1]] {
			s = append(s, i)
		}
	}

	res := 0
	for i := len(A) - 1; i >= 0; i-- {
		for len(s) != 0 && A[s[len(s)-1]] <= A[i] {
			res = max(res, i-s[len(s)-1])
			s = s[:len(s)-1]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
