package problem1004

func longestOnes(A []int, K int) int {
	A = append(A, 0)
	size := len(A)

	i := 0
	for i < size && A[i] == 0 {
		i++
	}

	ones, zeros := gen(A[i:])

	res := 0
	i, j := 0, 0
	one := 0
	for j < len(zeros) {
		if zeros[j] > K {
			res = max(res, one+ones[j]+K)
			for i < j && K < zeros[j] {
				K += zeros[i]
				one -= ones[i] + zeros[i]
				i++
			}
		}

		if K >= zeros[j] {
			K -= zeros[j]
			one += ones[j] + zeros[j]
			res = max(res, one)
		}

		j++
	}

	return min(res, size-1)
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

func gen(A []int) ([]int, []int) {
	ones, zeros := make([]int, 0, len(A)), make([]int, 0, len(A))
	size := len(A)
	a, i := 0, 0
	for i < len(A) {
		a = i
		for i < size && A[i] == 1 {
			i++
		}
		ones = append(ones, i-a)
		a = i
		for i < size && A[i] == 0 {
			i++
		}
		zeros = append(zeros, i-a)
	}
	ones = append(ones, 0)
	zeros = append(zeros, len(A))
	return ones, zeros
}
