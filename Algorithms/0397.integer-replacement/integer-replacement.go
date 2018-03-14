package problem0397

func integerReplacement(n int) int {
	// rec[i] == integerReplacement(i)
	rec := make(map[int]int)
	rec[1] = 0

	var ir func(int) int
	ir = func(i int) int {
		// 已经计算过的 i，可以直接从 rec 中读取出来
		if n, ok := rec[i]; ok {
			return n
		}

		if i%2 == 0 {
			rec[i] = ir(i/2) + 1
			return rec[i]
		}

		rec[i] = min(ir(i+1), ir(i-1)) + 1
		return rec[i]
	}

	return ir(n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
