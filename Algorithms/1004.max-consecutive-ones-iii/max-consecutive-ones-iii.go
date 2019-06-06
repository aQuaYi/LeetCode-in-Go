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
	return ones, zeros
}
