package problem0673

func findNumberOfLIS(a []int) int {
	n := len(a)
	res := 0
	maxLen := 0

	// length[i] == m 表示，以 a[i] 结尾的最长子序列的长度为 m
	length := make([]int, n)
	// count[i] == n 表示，以 a[i] 结尾的最长子序列有 n 个
	count := make([]int, n)

	for j := 0; j < n; j++ {
		length[j] = 1
		count[j] = 1
		for i := 0; i < j; i++ {
			if a[i] < a[j] {
				if length[j] == length[i]+1 {
					// 出现相同的长度
					count[j] += count[i]
				} else if length[j] < length[i]+1 {
					// 出现了长度的子序列
					length[j] = length[i] + 1
					count[j] = count[i]
				}
			}
		}

		// 对于每一个 j 都更新一次 res
		if maxLen == length[j] {
			res += count[j]
		} else if maxLen < length[j] {
			res = count[j]
			maxLen = length[j]
		}
	}

	return res
}
