package Problem0330

func minPatches(nums []int, n int) int {
	remain, count := n, 0
	isGetted := make([]bool, n+1)

	// 利用 nums 已有的元素，填充isGetted
	var dfs func(int, int)
	dfs = func(sum, idx int) {
		if idx == len(nums) {
			if 0 < sum && sum <= n && !isGetted[sum] {
				isGetted[sum] = true
				remain--
			}
			return
		}
		dfs(sum, idx+1)
		dfs(sum+nums[idx], idx+1)
	}

	dfs(0, 0)

	if remain == 0 {
		return 0
	}

	var i, j int
	i = 1
	for remain > 0 {
		for isGetted[i] {
			i++
		}

		count++
		remain--

		for j = n; 1 <= j-i; j-- {
			if !isGetted[j] && isGetted[j-i] {
				isGetted[j] = true
				remain--
			}
		}
	}

	return count
}
