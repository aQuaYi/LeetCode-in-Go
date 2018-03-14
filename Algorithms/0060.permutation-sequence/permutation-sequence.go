package problem0060

func getPermutation(n int, k int) string {
	if n == 0 {
		return ""
	}

	// 用于存放结果的字符
	res := make([]byte, n)

	// 存放待抓取的数字
	rec := make([]byte, n)
	for i := 0; i < n; i++ {
		rec[i] = byte(i) + '1'
	}

	// 由于排列的序号是从 1 开始的。
	// k 需要减去 1 ，好从 0 开始表示
	k--

	base := 1
	for i := 2; i < n; i++ {
		base *= i
	}

	for i := 0; i < n-1; i++ {
		idx := k / base
		res[i] = rec[idx]
		// 从 rec 中去除已经使用的数 rec[idx]
		rec = append(rec[:idx], rec[idx+1:]...)
		k %= base
		base /= (n - i - 1)
	}
	// 不要忘记最后一个数
	res[n-1] = rec[0]

	return string(res)
}
