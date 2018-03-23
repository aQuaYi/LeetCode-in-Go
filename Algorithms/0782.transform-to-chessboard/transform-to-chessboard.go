package problem0782

func movesToChessboard(b [][]int) int {
	N := len(b)

	// b 中的行有且只能有两种
	// 并且这两种行互补，例如
	// [1 0 1 0 1 1 0 0]
	// [0 1 0 1 0 0 1 1]
	// 一旦出现第 3 种，就可以返回 -1 了
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if b[0][0]^b[i][0]^b[0][j]^b[i][j] == 1 {
				// 等价于
				// if !((b[0][0] == b[i][0] && b[0][j] == b[i][j]) ||
				// (b[0][0] != b[i][0] && b[0][j] != b[i][j])) {
				return -1
			}
		}
	}

	rowSum, colSum, rowSwap, colSwap := 0, 0, 0, 0

	for i := 0; i < N; i++ {
		rowSum += b[0][i]
		colSum += b[i][0]
		// 先假设排列好的样子是
		// 0 1 0 1 0 1 ...
		// 1
		// 0
		// 1
		// 0
		// 1
		// .
		// .
		// .
		//
		// 那么 rowSwap 和 colSwap 分别统计了列和行上
		// 需要变动的元素的个数
		if b[i][0] == i%2 {
			rowSwap++
		}
		if b[0][i] == i%2 {
			colSwap++
		}
	}

	// 当 N 为偶数时， rowSum == N/2 且 colSum == N/2
	// 当 N 为奇数时， | rowSum - colSum | == 1
	// 才有可能交换成功
	// 否则，就可以返回 -1 了
	if rowSum < N/2 || (N+1)/2 < rowSum {
		// 一行中，拥有了太少或太多的 1
		return -1
	}

	if colSum < N/2 || (N+1)/2 < colSum {
		// 一列中，拥有了太少或太多的 1
		return -1
	}

	// 当 N 为奇数时
	if N%2 == 1 {
		// 如果 colSwap 也为奇数的话
		// 说明，我们前面预想的排列好的样子中
		// 行的样子反了
		if colSwap%2 == 1 {
			// 以前以为放对的位置，其实才是放错了的
			colSwap = N - colSwap
		}
		// 如果 rowSwap 也为奇数的话
		// 说明，我们前面预想的排列好的样子中
		// 列的样子反了
		if rowSwap%2 == 1 {
			// 以前以为放对的位置，其实才是放错了的
			rowSwap = N - rowSwap
		}
		// 其实可以从另一个角度来想
		// 由于只存在 0 和 1 两种元素，因此
		// colSwap 和 rowSwap 只能是偶数
		// 当他们为奇数时，说明，我们前面的预想是错的
	} else { // 当 N 为偶数时
		// 如果 colSwap 或 rowSwap 超过了 N/2
		// 同样说明我们预想的样子反了
		// 那不是经过最少步骤能够达到的样子
		colSwap = min(N-colSwap, colSwap)
		rowSwap = min(N-rowSwap, rowSwap)
	}

	// colSwap 和 rowSwap 是错位的元素个数
	// 调整次数需要 ÷2
	return (colSwap + rowSwap) / 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 本题来自于 https://leetcode.com/problems/transform-to-chessboard/discuss/114847/Easy-and-Concise-Solution-with-Explanation-C++JavaPython
