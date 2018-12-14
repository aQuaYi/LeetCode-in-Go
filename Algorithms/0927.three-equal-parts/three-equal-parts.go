package problem0927

func threeEqualParts(A []int) []int {
	size := len(A)
	indexs := make([]int, 0, size)
	for i, n := range A {
		if n == 1 {
			indexs = append(indexs, i)
		}
	}

	count := len(indexs)
	if count == 0 {
		return []int{0, size - 1}
	}

	if count%3 != 0 {
		return []int{-1, -1}
	}

	i1 := count / 3
	i2 := i1 * 2
	s0, s1, s2 := indexs[0], indexs[i1], indexs[i2]

	for s2 < size && A[s0] == A[s1] && A[s1] == A[s2] {
		s0++
		s1++
		s2++
	}

	if s2 == size {
		return []int{s0 - 1, s1}
	}

	return []int{-1, -1}
}
