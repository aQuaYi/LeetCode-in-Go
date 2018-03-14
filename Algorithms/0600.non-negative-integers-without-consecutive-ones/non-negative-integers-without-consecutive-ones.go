package problem0600

func findIntegers(num int) int {
	// binary 是 num 的二进制表示，逆序
	binary := make([]byte, 0, 32)
	for num > 0 {
		binary = append(binary, byte(num%2)+'0')
		num /= 2
	}

	n := len(binary)
	// a[i] == 长度为 i+1 且以 0 结尾的 binary 中不包含连续 1 的个数
	a := make([]int, n)
	// b[i] == 长度为 i+1 且以 1 结尾的 binary 中不包含连续 1 的个数
	b := make([]int, n)
	a[0] = 1
	b[0] = 1

	for i := 1; i < n; i++ {
		// 在 a[i-1] 中的 binary 右边加 0
		// 和
		// 在 b[i-1] 中的 binary 右边加 0
		// 可以得到 a[i] 中的 binary
		a[i] = a[i-1] + b[i-1]
		// 在 a[i-1] 中的 binary 右边加 1，可以得到 b[i] 中的 binary
		// 在 b[i-1] 中的 binary 右边加 1 会得到连续的 1，违反了题意
		b[i] = a[i-1]
	}

	res := a[n-1] + b[n-1]
	// 此时的 res 可能包含了 >num 的那一部分数，需要减去相关部分
	for i := n - 2; i >= 0; i-- {
		// 注意，此时的 binary 是逆序
		// 出现连续的 `1` 说明已经清理干净了
		if binary[i] == '1' && binary[i+1] == '1' {
			break
		}
		// 当初多加了 b[i]
		if binary[i] == '0' && binary[i+1] == '0' {
			res -= b[i]
		}
	}

	return res
}
