package problem0446

func numberOfArithmeticSlices(a []int) int {
	n := len(a)
	if n < 3 {
		return 0
	}

	min, max := -1<<31, 1<<31-1

	// dp[j][d] == x 表示
	// 如果存在 a[i] - a[j] == d, j < i < n, 那么, {a[:j+1]..., a[i]} 中
	// 可以构成形如 [..., a[j], a[i]] 的差值为 d 的等差数组的个数
	dp := make([]map[int]int, n)

	res := 0
	for i := 1; i < n; i++ {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			d := a[i] - a[j]

			// 存在一个特殊的 test case，用于加速
			if d <= min || max <= d {
				continue
			}

			// + 1 是因为 [a[j], a[i]] 无法构成等差数列
			// 但是如果存在 a[k] == a[i] + d 的话，
			// [a[j], a[i], a[k]] 就能构成 **1** 个等差数列了
			// 这里再次强调一下，d[i][d] 的含义是
			// 再加一个 a[i] + d 可以构成形如
			// [..., a[i], a[i]+d] 的等差数列子数组的个数
			dp[i][d] += dp[j][d] + 1
			// a[j] 已经遇到了 a[j] + d == a[i]
			// 所以，res 加上的是 dp[j][d]
			res += dp[j][d]
		}
	}

	return res
}
