package problem1013

func canThreePartsEqualSum(A []int) bool {
	n := len(A)
	sums := make([]int, n)
	s := 0
	for i := 0; i < n; i++ {
		s += A[i]
		sums[i] = s
	}
	if s%3 != 0 {
		return false
	}

	s /= 3
	i := 0
	for i < n && sums[i] != s {
		i++
	}

	s *= 2
	j := n - 1
	for 0 <= j && sums[j] != s {
		j--
	}

	return i < j
}
