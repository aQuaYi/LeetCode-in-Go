package problem0823

import (
	"sort"
)

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
	for i := range ress {
		ress[i] = 1
	}

	for i := 1; i < len(A); i++ {
		for j := 0; j < i; j++ {
			quotient, remainder := A[i]/A[j], A[i]%A[j]

			if remainder != 0 ||
				(factorIndex[quotient] == 0 && quotient != A[0]) {
				continue
			}

			k := factorIndex[quotient]

			t := mod(ress[j] * ress[k])

			ress[i] = mod(ress[i] + t)

		}
	}

	return sum(ress)
}

func mod(n int) int {
	return n % modulo
}

func sum(a []int) int {
	res := 0
	for i := range a {
		res += a[i]
		res = mod(res)
	}
	return res
}
