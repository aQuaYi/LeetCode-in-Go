package problem0456

func find132pattern(a []int) bool {
	// 题目要求找到 i < j < k
	// 使得 ai < ak < aj
	ak := -1 << 31
	// ajStack 存放了所有符合 ai < ak < aj 的 aj
	ajStack := make([]int, 0, len(a))

	for i := len(a) - 1; 0 <= i; i-- {

		if a[i] < ak {
			// 找到了题目要求的 ai < aj < ak
			return true
		}

		// ajStack 中的数字的 index 都比 i 大
		// 从 ajStack 中挑选一个仅次于 a[i] 的数字作为新的 ak
		// 然后，把 a[i] 当做下一轮 for 的 aj 放入 ajStack
		for len(ajStack) > 0 &&
			ajStack[len(ajStack)-1] < a[i] {
			ak = ajStack[len(ajStack)-1]
			ajStack = ajStack[:len(ajStack)-1]

			// 根据 stack 的进出栈规则
			// ajStack 中的 ax 及其 ax 在 a 中的索引号 x，都是递减的
			// 理解这一点很关键
			// 始终维持了 ak < ax ∈ ajStack 且 ak 尽可能地大 这样的关系
		}

		ajStack = append(ajStack, a[i])
	}

	return false
}
