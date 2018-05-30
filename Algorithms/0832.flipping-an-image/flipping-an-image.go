package problem0832

func flipAndInvertImage(A [][]int) [][]int {
	for k := 0; k < len(A); k++ {
		i, j := 0, len(A[k])-1
		for i < j {
			A[k][i], A[k][j] = invert(A[k][j]), invert(A[k][i])
			i++
			j--
		}
		if i == j { // 当 len(A[k]) 的长度为奇数时，处理正中间的数
			A[k][i] = invert(A[k][i])
		}
	}
	return A
}

func invert(i int) int {
	if i == 0 {
		return 1
	}
	return 0
}
