package problem0775

func isIdealPermutation(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] == i {
			continue
		}

		// 只有 **仅有** 相邻的两个数交换的时候
		// 才有可能满足题意，局部逆序等于全局逆序
		if a[i] == i+1 && a[i+1] == i {
			i++
		} else {
			return false
		}
	}

	return true
}
