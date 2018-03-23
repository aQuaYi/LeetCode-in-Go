package problem0782

func movesToChessboard(board [][]int) int {

	return 0
}

// 想要 board 可行，需要满足以下条件
// 	1) board 中的行，只有两种，且分别是 r 和 r 的反集
// 	2) r 中 1 的个数
// 		当 N 为偶数时， 1 的个数为 N/2
// 		当 N 为奇数时， 1 的个数为 N/2 或 N/2+1
// 	3) board 中全部 0 和 1 的个数
// 		当 N 为偶数时， 个数相等
// 		当 N 为奇数时， 个数只能相差一个
func isValid(board [][]int) bool {
	N := len(board)

	rowOne := sum(board[0])
	if (N%2 == 0 && rowOne != N/2) ||
		(N%2 == 1 && rowOne != N/2 && rowOne != N/2+1) {
		return false
	}

	countSame := 1
	for i := 1; i < N; i++ {
		if isSame(board[0], board[i]) {
			countSame++
		} else if !isContrary(board[0], board[i]) {
			return false
		}
	}
	if (N%2 == 0 && countSame != N/2) ||
		(N%2 == 1 && countSame != N/2 && countSame != N/2+1) {
		return false
	}

	boardOne := rowOne * countSame
	NN := N * N
	if (NN%2 == 0 && boardOne != NN/2) ||
		(NN%2 == 1 && boardOne != NN/2 && boardOne != NN/2+1) {
		return false
	}

	return true
}

func sum(a []int) int {
	res := 0
	for i := range a {
		res += a[i]
	}
	return res
}

func isSame(a, b []int) bool {
	for i := range a {
		if a[i] == b[i] {
			continue
		}
		return false
	}
	return true
}

func isContrary(a, b []int) bool {
	for i := range a {
		if a[i]+b[i] == 1 {
			continue
		}
		return false
	}
	return true
}
