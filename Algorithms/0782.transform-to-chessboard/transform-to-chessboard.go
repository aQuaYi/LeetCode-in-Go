package problem0782

func movesToChessboard(board [][]int) int {
	isValid, leftTop := checkValid(board)
	if !isValid {
		return -1
	}

	N := len(board[0])
	if N%2 == 1 && board[0][0] != leftTop {
		i := 1
		for i < N {
			if board[i][0] == leftTop {
				break
			}
			i += 2
		}
		board[0], board[i] = board[i], board[0]
		return 1 + exchangeCount(board)
	}
	return exchangeCount(board)
}

// 想要 board 可行，需要满足以下条件
// 	1) board 中的行，只有两种，且分别是 r 和 r 的反集
// 	2) r 中 1 的个数
// 		当 N 为偶数时， 1 的个数为 N/2
// 		当 N 为奇数时， 1 的个数为 N/2 或 N/2+1
// 	3) board 中全部 0 和 1 的个数
// 		当 N 为偶数时， 个数相等
// 		当 N 为奇数时， 个数只能相差一个
func checkValid(board [][]int) (bool, int) {
	N := len(board)

	rowOne := sum(board[0])
	if (N%2 == 0 && rowOne != N/2) ||
		(N%2 == 1 && rowOne != N/2 && rowOne != N/2+1) {
		return false, 0
	}

	countSame := 1
	for i := 1; i < N; i++ {
		if isSame(board[0], board[i]) {
			countSame++
		} else if !isContrary(board[0], board[i]) {
			return false, 0
		}
	}
	if (N%2 == 0 && countSame != N/2) ||
		(N%2 == 1 && countSame != N/2 && countSame != N/2+1) {
		return false, 0
	}

	boardOne := rowOne*countSame + (N-rowOne)*(N-countSame)
	NN := N * N
	if (NN%2 == 0 && boardOne != NN/2) ||
		(NN%2 == 1 && boardOne != NN/2 && boardOne != NN/2+1) {
		return false, 0
	}

	res := 0
	if NN%2 == 1 && boardOne == NN/2+1 {
		// 当 N 为奇数时，
		// 如果 1 比 0 多一个，则左上角的元素必为 1
		res = 1
	}

	return true, res
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
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isContrary(a, b []int) bool {
	for i := range a {
		if a[i]+b[i] != 1 {
			return false
		}
	}
	return true
}

func exchangeCount(board [][]int) int {
	res := 0
	N := len(board[0])
	for i := 2; i < N; i += 2 {
		if board[i][0] != board[0][0] {
			res++
		}
	}
	for j := 2; j < N; j += 2 {
		if board[0][j] != board[0][0] {
			res++
		}
	}

	return res
}
