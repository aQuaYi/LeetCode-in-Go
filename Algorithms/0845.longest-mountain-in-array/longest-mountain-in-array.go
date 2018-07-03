package problem0845

func longestMountain(A []int) int {
	size := len(A)
	if size <= 2 {
		return 0
	}

	res, up, down := 0, 0, 0
	for i := 1; i < size; i++ {
		if (down > 0 && A[i-1] < A[i]) ||
			A[i-1] == A[i] {
			up, down = 0, 0
		}
		if A[i-1] < A[i] {
			up++
		}
		if A[i-1] > A[i] {
			down++
		}
		if up > 0 && down > 0 && up+down+1 > res {
			res = up + down + 1
		}
	}

	return res
}
