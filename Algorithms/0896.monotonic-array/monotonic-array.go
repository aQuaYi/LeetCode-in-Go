package problem0896

func isMonotonic(A []int) bool {
	return isMonotoneIncreasing(A) || isMonotoneDecreasing(A)
}

func isMonotoneIncreasing(A []int) bool {
	size := len(A)
	for i := 1; i < size; i++ {
		if A[i-1] > A[i] {
			return false
		}
	}
	return true
}

func isMonotoneDecreasing(A []int) bool {
	size := len(A)
	for i := 1; i < size; i++ {
		if A[i-1] < A[i] {
			return false
		}
	}
	return true
}
