package problem0927

// ref: https://leetcode.com/problems/three-equal-parts/discuss/183922/C++-O(n)-time-O(1)-space-12-ms-with-explanation-and-comments

var noPossible = []int{-1, -1}

func threeEqualParts(A []int) []int {
	size := len(A)

	indexs := make([]int, 0, size)
	for i, val := range A {
		if val == 1 {
			indexs = append(indexs, i)
		}
	}

	count := len(indexs)

	if count == 0 {
		return []int{0, size - 1}
	}

	if count%3 != 0 {
		return noPossible
	}

	i1, i2 := count/3, count/3*2
	s0, s1, s2 := indexs[0], indexs[i1], indexs[i2]

	for s2 < size && A[s0] == A[s1] && A[s1] == A[s2] {
		s0++
		s1++
		s2++
	}

	if s2 == size {
		return []int{s0 - 1, s1}
	}

	return noPossible
}
