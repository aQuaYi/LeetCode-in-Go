package problem0962

func maxWidthRamp(A []int) int {
	size := len(A)
	res := 0

	indexs := newIndexs(size)

	for sz := 1; sz < size; sz *= 2 {
		for lo := 0; lo < size-sz; lo += sz + sz {
			mid, hi := lo+sz-1, min(lo+sz+sz-1, size-1)
			res = max(res, check(A, indexs, lo, mid, hi))
			merge(A, indexs, lo, mid, hi)
		}
	}

	return res
}

func newIndexs(size int) []int {
	res := make([]int, size)
	for i := range res {
		res[i] = i
	}
	return res
}

func check(A, indexs []int, lo, mid, hi int) int {
	width := 0
	for i := lo; i <= mid; i++ {
		for j := mid + 1; j <= hi; j++ {
			if A[i] <= A[j] {
				width = max(width, indexs[j]-indexs[i])
			}
		}
	}
	return width
}

func merge(A, indexs []int, lo, mid, hi int) {
	i, j := lo, mid+1
	ta := make([]int, hi-lo+1)
	ti := make([]int, hi-lo+1)
	for k := 0; k <= hi-lo; k++ {
		if i > mid {
			ta[k] = A[j]
			ti[k] = indexs[j]
			j++
			continue
		}

		if j > hi {
			ta[k] = A[i]
			ti[k] = indexs[i]
			i++
			continue
		}

		if A[i] <= A[j] {
			ta[k] = A[i]
			ti[k] = indexs[i]
			i++
		} else {
			ta[k] = A[j]
			ti[k] = indexs[j]
			j++
		}
	}

	copy(A[lo:hi+1], ta)
	copy(indexs[lo:hi+1], ti)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
