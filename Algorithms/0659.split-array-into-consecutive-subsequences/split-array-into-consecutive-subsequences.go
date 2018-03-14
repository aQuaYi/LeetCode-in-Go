package problem0659

func isPossible(nums []int) bool {
	size := len(nums)
	// count[i]=j 表示 nums 中 i 出现了 j 次
	count := make(map[int]int, size)
	// next[i]=j 表示有 j 个子序列，可以通过在末尾添加 i 来延长
	next := make(map[int]int, size)
	for i := range nums {
		count[nums[i]]++
	}

	for _, n := range nums {
		if count[n] == 0 {
			// 在上一个 for 循环中，被用来添加新的子序列了
			continue
		}
		count[n]--

		if next[n] > 0 {
			// 可以把 n 放入别的子序列末尾的时候，就放在末尾
			next[n]--
			next[n+1]++
		} else if count[n+1] > 0 && count[n+2] > 0 {
			// 不能放在放在末尾的时候，就创建新的子序列
			count[n+1]--
			count[n+2]--
			next[n+3]++
		} else {
			// 都不行的时候，就表示不可能满足题目要求，返回 false
			return false
		}

	}

	return true
}
