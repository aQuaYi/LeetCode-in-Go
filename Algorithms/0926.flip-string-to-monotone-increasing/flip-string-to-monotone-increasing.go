package problem0926

func minFlipsMonoIncr(S string) int {
	size := len(S)

	ones := make([]int, 0, size)
	zeros := make([]int, 0, size)

	for i := 0; i < size; i++ {
		if S[i] == '1' {
			ones = append(ones, i)
		} else {
			zeros = append(zeros, i)
		}
	}

	oneSize := len(ones)
	zeroSize := len(zeros)

	res := size

	i, j := 0, 0
	for i < oneSize {
		for j < zeroSize && zeros[j] < ones[i] {
			j++
		}
		res = min(res, i+zeroSize-j)
		i++
	}

	res = min(res, oneSize)

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
