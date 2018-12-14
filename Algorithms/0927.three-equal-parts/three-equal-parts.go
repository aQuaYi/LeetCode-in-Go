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
		return []int{0, 2}
	}

	if count%3 != 0 {
		// 此时，无法保证每个部分都有相同数目的 1
		// 所以，不可能出现可行解
		return noPossible
	}

	c1, c2 := count/3, count/3*2
	s0, s1, s2 := indexs[0], indexs[c1], indexs[c2]

	// 此时 A[s0:s1]， A[s1:s2] 和 A[s2:] 中分别有三分之一的 1
	// 且
	// A[s0]==A[s1]==A[s2]==1 三个部分的第一个 1 已经对齐
	// 可以开始逐个比较
	for s2 < size &&
		A[s0] == A[s1] &&
		A[s1] == A[s2] {
		s0++
		s1++
		s2++
	}

	if s2 == size {
		// 此时， 所有的 1 都被比较过了
		// 且三段中的 1 分布相同。
		// 没有被比较的元素都是 0
		return []int{s0 - 1, s1}
	}

	return noPossible
}
