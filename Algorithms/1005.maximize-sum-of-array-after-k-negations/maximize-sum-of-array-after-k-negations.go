package problem1005

import "sort"

func largestSumAfterKNegations(A []int, K int) int {
	size := len(A)

	i, j := 0, size-1
	minAbs := abs(A[0])
	for i <= j {
		minAbs = min(minAbs, abs(A[i]))
		if A[i] >= 0 { // move non-negative to right side
			A[i], A[j] = A[j], A[i]
			j--
			continue
		}
		i++
	}

	negatives := A[:j+1]
	negSize := j + 1

	sum := 0
	if K >= negSize { // all negative could convert to positive
		if (K-negSize)&1 == 1 { // someone need keep as negative
			sum -= minAbs * 2 // choose minAbs keep as negative
		}
	} else {
		sort.Ints(negatives)
		// min K negative could convert to positive
		// sort make min K negative in negatives[:K]
	}

	index := min(K, negSize)
	for i := 0; i < index; i++ {
		sum -= A[i]
	}
	for i := index; i < size; i++ {
		sum += A[i]
	}

	return sum
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
