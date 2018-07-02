package problem0845

func longestMountain(A []int) int {
	size := len(A)
	if size <= 2 {
		return 0
	}

	isAscent := false
	isDecent := false

	res := 0
	temp := 0

	for i := 1; i < size; i++ {
		switch {
		case A[i-1] < A[i]:
			if isAscent {
				temp++
			} else {
				isAscent = true
				isDecent = false
				temp = 2
			}
		case A[i-1] > A[i]:
			if isAscent || isDecent {
				temp++
				res = max(res, temp)
				isAscent = false
				isDecent = true
			}
		default:
			isAscent, isDecent = false, false
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
