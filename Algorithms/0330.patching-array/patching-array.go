package Problem0330

func minPatches(nums []int, n int) int {
	var i int

	remain, count := log2(n)+1, 0
	isGetted := make([]bool, remain)

	pow2 := make([]int, 1, remain)
	pow2[0] = 1
	for i = 1; i < remain; i++ {
		pow2[i] = pow2[i-1] * 2
	}

	// 利用 nums 已有的元素，填充isGetted
	var dfs func(int, int)
	var idx int 
	dfs = func(sum, idx int) {
		if idx == len(nums) {
			if 0 < sum && sum <= n {
				idx = log2(sum)
				if !isGetted[idx] && pow2[idx] == sum  {
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

// log2(n) 返回 n 以 2 为底的对数的整数部分
func log2(n int) int {
	res := 0
	for n > 1 {
		n = n >> 1
		res++
	}
	return res
}
