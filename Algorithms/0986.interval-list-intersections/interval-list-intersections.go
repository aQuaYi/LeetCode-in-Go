package problem0986

func intervalIntersection(A [][]int, B [][]int) [][]int {
	sizeA, sizeB := len(A), len(B)
	res := make([][]int, 0, sizeA+sizeB)
	for i, j := 0, 0; i < sizeA && j < sizeB; {
		switch {
		case A[i][1] < B[j][0]:
			i++
		case B[j][1] < A[i][0]:
			j++
		default:
			res = append(res, []int{
				max(A[i][0], B[j][0]),
				min(A[i][1], B[j][1]),
			})
			if A[i][1] < B[j][1] {
				i++
			} else {
				j++
			}
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
