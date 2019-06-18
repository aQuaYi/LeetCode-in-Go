package problem1031

import "sort"

func maxSumTwoNoOverlap(A []int, L int, M int) int {
	ls, ms := sums(A, L), sums(A, M)
	lSize, mSize := len(ls), len(ms)
	res := 0
	for i := 0; i < lSize; i++ {
		if ls[i].sum+ms[0].sum <= res {
			break
		}
		j := 0
		for j < mSize && ls[i].isOverlapping(ms[j]) {
			j++
		}
		if j < mSize {
			res = max(res, ls[i].sum+ms[j].sum)
		}
	}
	return res
}

type entry struct {
	sum, left, right int
}

func (e *entry) isOverlapping(other *entry) bool {
	return max(e.left, other.left) <= min(e.right, other.right)
}

func sums(A []int, l int) []*entry {
	n := len(A)
	res := make([]*entry, n-l+1)
	sum := 0
	for i := 0; i < l; i++ {
		sum += A[i]
	}
	for i := l; i < n; i++ {
		res[i-l] = &entry{
			sum:   sum,
			left:  i - l,
			right: i - 1,
		}
		sum += A[i] - A[i-l]
	}
	res[n-l] = &entry{
		sum:   sum,
		left:  n - l,
		right: n - 1,
	}
	sort.Slice(res, func(i int, j int) bool {
		return res[i].sum > res[j].sum
	})
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
