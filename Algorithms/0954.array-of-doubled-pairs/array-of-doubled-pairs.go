package problem0954

import "sort"

func canReorderDoubled(A []int) bool {
	size := len(A)
	A1, A2 := make([]int, 0, size), make([]int, 0, size)
	for _, v := range A {
		if v >= 0 {
			A1 = append(A1, v)
		} else {
			A2 = append(A2, -v)
		}
	}

	return isPossible(A1) && isPossible(A2)
}

func isPossible(A []int) bool {
	size := len(A)

	if size%2 == 1 {
		return false
	}

	sort.Ints(A)

	i, j := 0, 1
	count := size / 2
	for j < size {
		for A[i] < 0 {
			i++
		}
		double := A[i] * 2
		i++
		for j < size && A[j] < double {
			j++
		}
		if j == size || A[j] != double {
			return false
		}
		A[j] = -1
		j++
		count--
	}

	return count == 0
}
