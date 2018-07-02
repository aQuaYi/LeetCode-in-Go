package problem0845

func longestMountain(A []int) int {
	size := len(A)
	if size <= 2 {

	}

	isAscent := false
	isDecent := false

	res := 0
	temp := 0

	for i := 1; i < size; i++ {
		if A[i-1] == A[i] {
			isAscent, isDecent = false, false
			continue
		}

		if !isAscent && A[i-1] < A[i] {
			isAscent = true
			temp = 1
			for i < size && A[i-1] < A[i] {
				temp++
				i++
			}
		}

		if (isAscent || isDecent) && A[i-1] > A[i] {
			temp++
			isDecent = true
			res = max(res, temp)
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
