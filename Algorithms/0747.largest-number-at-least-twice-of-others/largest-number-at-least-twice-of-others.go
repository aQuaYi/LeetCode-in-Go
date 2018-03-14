package problem0747

func dominantIndex(a []int) int {
	n := len(a)
	if n == 1 {
		return 0
	}

	// i1 是 a 中第 1 大的数的索引号
	// i2 是 a 中第 2 大的数的索引号
	i1, i2 := 0, 1
	if a[i1] < a[i2] {
		i1, i2 = i2, i1
	}

	// 寻找真正的 i1 和 i2
	for i := 2; i < n; i++ {
		if a[i1] < a[i] {
			i2, i1 = i1, i
		} else if a[i2] < a[i] {
			i2 = i
		}
	}

	if a[i2] == 0 || a[i1]/a[i2] >= 2 {
		return i1
	}
	return -1
}
