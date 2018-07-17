package problem0457

func circularArrayLoop(nums []int) bool {
	size := len(nums)

	// 缩小 nums[i] 的范围
	// 特别地，如果 nums[i] 是一个单节点 loop
	// nums[i] 会变成 0
	for i := range nums {
		nums[i] %= size
	}

	// next 返回从 i 出发，到达下一个位置的索引值
	next := func(i int) int {
		return (size + i + nums[i]) % size
	}

	// 题目要求 loop 必须是单方向的
	// 所以在找到 loop 后，还需要检查方向
	isSingleForward := func(i int) bool {
		j := next(i)
		a, b := nums[i], nums[j]
		for j != i {
			if a*b < 0 {
				// 此时发生了转向
				return false
			}
			j = next(j)
			b = nums[j]
		}
		return true
	}

	for i := range nums {
		// slow/fast 搜索
		s, f := next(i), next(next(i))
		for s != f {
			s, f = next(s), next(next(f))
		}

		if nums[s] != 0 {
			return isSingleForward(s)
		}
	}

	return false
}
