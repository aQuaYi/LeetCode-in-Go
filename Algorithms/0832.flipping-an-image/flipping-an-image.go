package problem0832

func flipAndInvertImage(A [][]int) [][]int {
	return invert(reverse(A))
}

func reverse(A [][]int) [][]int {
	for k := 0; k < len(A); k++ {
		i, j := 0, len(A[k])-1
		for i < j {
			A[k][i], A[k][j] = A[k][j], A[k][i]
			i++
			j--
		}
	}
	return A
}

func invert(A [][]int) [][]int {
	for i := range A {
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] == 0 {
				A[i][j] = 1
			} else {
				A[i][j] = 0
			}
		}
	}
	return A
}
