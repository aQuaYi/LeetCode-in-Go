package problem0457

func circularArrayLoop(nums []int) bool {
	size := len(nums)

	// next 返回从 i 出发，到达下一个位置的索引值
	next := func(i int) int {
		return (((i + nums[i]) % size) + size) % size
	}

	// 把所有单节点 loop 中的值，变成 0
	// 这样在 slow/fast 搜索中，s 和 f 会止步与此处，很妙
	for i := range nums {
		if next(i) == i {
			nums[i] = 0
		}
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
		if nums[i] == 0 {
			continue
		}

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
