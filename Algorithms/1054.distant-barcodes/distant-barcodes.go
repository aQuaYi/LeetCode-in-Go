package problem1054

import "sort"

func rearrangeBarcodes(A []int) []int {
	n := len(A)

	count := [10001]int{}
	for _, a := range A {
		count[a]++
	}

	sort.Slice(A, func(i int, j int) bool {
		if count[A[i]] == count[A[j]] {
			return A[i] < A[j]
		}
		return count[A[i]] > count[A[j]]
	})

	res := make([]int, n)
	i := 0
	for _, a := range A {
		res[i] = a
		i += 2
		if i >= n {
			i = 1
		}
	}

	return res
}
