package problem1122

import "sort"

func relativeSortArray(A, B []int) []int {
	serial := [1001]int{}
	for i, b := range B {
		serial[b] = i + 1
	}

	indexOf := func(i int) int {
		Ai := A[i]
		sAi := serial[Ai]
		if sAi > 0 {
			return sAi
		}
		return 1000 + Ai
	}

	less := func(i, j int) bool {
		return indexOf(i) < indexOf(j)
	}

	sort.Slice(A, less)

	return A
}
