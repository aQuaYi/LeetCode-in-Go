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
		ress[i] = 1
		for j := 0; j < i; j++ {
			quotient, remainder := A[i]/A[j], A[i]%A[j]
			k, isFactor := factorIndex[quotient]

			if remainder != 0 || !isFactor {
				continue
			}

			ress[i] = mod(ress[i] + ress[j]*ress[k])
		}
	}

	return sum(ress)
}

func sum(a []int) int {
	res := 0
	for i := range a {
		res += a[i]
		res = mod(res)
	}
	return res
}

func mod(n int) int {
	return n % modulo
}
