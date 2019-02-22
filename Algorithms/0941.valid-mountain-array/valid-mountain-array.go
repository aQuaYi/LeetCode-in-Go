package problem0941

func validMountainArray(A []int) bool {
	size := len(A)
	if size < 3 {
		return false
	}

	i := 1
	for i < size && A[i-1] < A[i] {
		i++
	}

	index := i

	for i < size && A[i-1] > A[i] {
		i++
	}

	return index < i && i == size
}
