package problem0930

func numSubarraysWithSum(A []int, S int) int {
	if S == 0 {
		return zero(A)
	}

	size := len(A)
	i, j, sum := 0, 0, 0

	for j < size && sum < S {
		sum += A[j]
		j++
	}

	res := 0
	for i < size && j <= size && sum == S {
		left := 1
		for i < j && A[i] == 0 {
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

func zero(A []int) int {
	res := 0
	tmp := 0
	for _, a := range A {
		if a == 1 {
			tmp = 0
		} else {
			tmp++
			res += tmp
		}
	}
	return res
}
