package problem0823

import "sort"

const (
	modulo = 1E9 + 7
)

func numFactoredBinaryTrees(A []int) int {
	sort.Ints(A)

	factorIndex := make(map[int]int, len(A))
	for i := range A {
		factorIndex[A[i]] = i
	}

	ress := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		// 自身算一个
		ress[i] = 1
		for j := 0; j < i; j++ {
			quotient, remainder := A[i]/A[j], A[i]%A[j]
			k, isFactor := factorIndex[quotient]

			if remainder != 0 || !isFactor {
				continue
			}

			// 左边和右边的组合数
			// 题目还说了 A[x] is unique
			ress[i] += ress[j] * ress[k]
		}
		ress[i] = mod(ress[i])
	}

	return sum(ress)
}

func sum(a []int) int {
	res := 0
	for i := range a {
		res += a[i]
	}
	return mod(res)
}

func mod(n int) int {
	return n % modulo
}
