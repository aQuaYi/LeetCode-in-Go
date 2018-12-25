package problem0930

func numSubarraysWithSum(A []int, S int) int {
	size := len(A)
	i, j, sum := 0, 0, 0

	for j < size && sum < S {
		sum += A[j]
		j++
	}

	res := 0
	for i < size && j <= size {
		left := 1
		for i < size && A[i] == 0 {
			i++
			left++
		}

		right := 1
		for j < size && A[j] == 0 {
			j++
			right++
		}

		res += left * right

		i++
		j++
	}

	return res
}
