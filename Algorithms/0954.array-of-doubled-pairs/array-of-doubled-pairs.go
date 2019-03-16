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
	// 把负数单独分出来，并转换成正数的数列
	// 两个部分，就可以用相同的逻辑来处理了。
	return isPossible(A1) && isPossible(A2)
}

func isPossible(A []int) bool {
	size := len(A)

	if size%2 == 1 {
		return false
	}

	sort.Ints(A)

	count := 0
	i, j := 0, 1
	for j < size {
		for A[i] < 0 {
			i++
		}
		double := A[i] * 2
		for (j < size && A[j] < double) ||
			j <= i {
			j++
		}
		if j == size || A[j] != double {
			return false
		}
		A[j] = -1
		count++
		i++
		j++
	}

	return count == size/2
}
