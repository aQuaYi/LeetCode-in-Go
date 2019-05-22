package problem1013

import "sort"

func canThreePartsEqualSum(A []int) bool {
	n := len(A)
	sums := make([]int, n)
	sums[0] = A[0]
	for i := 1; i < n; i++ {
		sums[i] = sums[i-1] + A[i]
	}
	sum := sums[n-1]
	if sum%3 != 0 {
		return false
	}

	sort.Ints(sums)

	i := sort.SearchInts(sums, sum/3)
	if i == n || sums[i] != sum/3 {
		return false
	}

	j := sort.SearchInts(sums, sum/3*2)
	if j == n || sums[j] != sum/3*2 {
		return false
	}

	return i < j
}
