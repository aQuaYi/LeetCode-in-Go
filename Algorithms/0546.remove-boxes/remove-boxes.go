package problem0546

func removeBoxes(boxes []int) int {
	n := len(boxes)

	// dp[l][r][k] 的值是，
	// 在 boxes[l:r+1] 右侧有 k 个和 boxes[r] 一样颜色的 box 再加上 boxes[l:r+1] 可以获取的最大值
	// 所以，题目要求的是 dp[0][n-1][0] 的值
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
		}
	}

	return remove(boxes, dp, 0, n-1, 0)
}

func remove(boxes []int, dp [][][]int, l, r, k int) int {
	if l > r {
		return 0
	}

	// 已经计算过的情况，直接返回
	if dp[l][r][k] > 0 {
		return dp[l][r][k]
	}

	// 缩短 boxes[l:r+1] 直到 boxes[r-1] != boxes[r]
	// 这是一种优化行为，因为
	// 4*4 > 3*3 + 1*1，同时消除尽可能多的元素，肯定比分开消除好
	// 把下列 for 循环注释掉，也可以得到正确的解答，只是慢一些
	for l < r && boxes[r-1] == boxes[r] {
		r--
		k++
	}

	// 第一种选择，把右侧的 k 个元素和 boxes[r] 现在就清除掉
	res := (k+1)*(k+1) + remove(boxes, dp, l, r-1, 0)

	// 第二种选择，把右侧的 k 个元素 和 boxes[r] 留着，等着与 boxes[l:r] 中和 boxes[r]相同颜色的 box 汇合，变得更多了以后，再消除
	for m := r - 1; l <= m; m-- {
		if boxes[r] == boxes[m] {
			// 此时，boxes[m] 与 boxes[r] 的颜色相同
			// remove(boxes, dp, m+1, r-1, 0)+remove(boxes, dp, l, m, k+1) 的含义是
			// 先把 boxes[m+1:r] 清除掉 → remove(boxes, dp, m+1, r-1, 0)
			// boxes[l:m+1] 右侧就有了 k+1 个和 boxes[m] 一样颜色的盒子了，继续递归搜索
			res = max(res, remove(boxes, dp, m+1, r-1, 0)+remove(boxes, dp, l, m, k+1))
		}
	}

	dp[l][r][k] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
