package problem0786

import (
	"sort"
)

// 二分法查找
func kthSmallestPrimeFraction(A []int, K int) []int {
	sort.Ints(A)
	lo, hi := 0.0, 1.0

	for {
		mid := (lo + hi) / 2

		p, q, count := countUnder(mid, A)

		if count < K {
			lo = mid
		} else if count > K {
			hi = mid
		} else {
			return []int{p, q}
		}
	}

}

//
func countUnder(mid float64, A []int) (int, int, int) {
	n, p, q, count := len(A), 0, 1, 0

	// A[i] 和 p 表示分子
	// A[j] 和 q 表示分母
	for i := 0; i < n-1; i++ {
		lo, hi := i, n
		for lo < hi {
			m := (hi + lo) / 2
			if float64(A[i]) <= mid*float64(A[m]) {
				hi = m
			} else {
				lo = m + 1
			}
		}
		j := lo
		count += n - j

		if n-j > 0 && p*A[j] < q*A[i] {
			p = A[i]
			q = A[j]
		}

	}

	return p, q, count
}
