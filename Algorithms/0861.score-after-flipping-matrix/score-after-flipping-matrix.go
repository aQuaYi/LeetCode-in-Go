package problem0861

import (
	"sort"
)

func matrixScore(A [][]int) int {
	m, n := len(A), len(A[0])

	toggleRow := func(i int) {
		for j := 0; j < n; j++ {
			A[i][j] ^= 1
		}
	}

	toggleCol := func(j int) {
		for i := 0; i < m; i++ {
			A[i][j] ^= 1
		}
	}

	for i := 0; i < m; i++ {
		A[i] = append(A[i], b2d(A[i]))
	}

	sort.Slice(A, func(i int, j int) bool {
		return A[i][n] < A[j][n]
	})

	for j := 0; j < n; j++ {
		if A[0][j] != 0 {
			continue
		}
		toggleCol(j)
	}

	for i := 0; i < m; i++ {
		if A[i][0] == 0 {
			toggleRow(i)
		}
	}

	res := 0
	for i := 0; i < m; i++ {
		res += b2d(A[i][:n])
	}

	return res
}

func b2d(a []int) int {
	res := 0
	for _, n := range a {
		res = res*2 + n
	}
	return res
}
