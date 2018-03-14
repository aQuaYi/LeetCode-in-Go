package problem0204

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}

	// composite：合数，prime 的反义词
	isComposite := make([]bool, n)
	// 小于 n 的数中，有一半是偶数，除2以外，都不是质数
	// 所以，count 的上限是 n / 2
	count := n / 2

	for i := 3; i*i < n; i += 2 {
		if isComposite[i] {
			continue
		}

		// j 是 i 的倍数，所以 j 肯定不是 质数
		// 对所有的 j 进行标记
		for j := i * i; j < n; j += 2 * i {
			if !isComposite[j] {
				// j 还没有被别的 i 标记过了
				// 修改 count
				count--
				// 对 j 进行标记
				isComposite[j] = true
			}
		}
	}

	return count
}
