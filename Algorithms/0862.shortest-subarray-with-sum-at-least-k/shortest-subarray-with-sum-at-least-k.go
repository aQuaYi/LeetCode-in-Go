package problem0862

func shortestSubarray(a []int, K int) int {
	size := len(a)

	sums := make([]int, size+1)
	s := 0
	for i, n := range a {
		if n == K {
			return 1
		}
		s += n
		sums[i+1] = s
	}

	initialValue := size * 2
	res := initialValue

	deque := make([]int, size+1)
	deque[0] = 0
	first, last := 0, 0
	// deque[first:last+1] 中保存了以后会用到的 sums 中元素的索引号

	for i := 1; i <= size; i++ {
		for first <= last && sums[i]-sums[deque[first]] >= K {
			// 由于 i 递增
			// 即使以后会有 j>i 使得 sums[j] - sums[deque[first]] >= K
			// 但是由于 j>i, 导致 j-deque[first] > i-deque[first] 不会是更短的答案。
			// 所以，可以把 deque[first] 删除
			res = min(res, i-deque[first])
			first++
		}
		for first <= last && sums[i] <= sums[deque[last]] {
			// 如果存在 j>i>deque[last] 使得 sums[j] - sums[deque[last]] >= K
			// 由于 sums[deque[last]] >= sums[i] 则
			// sums[j] - sums[i] >= K 一定成立，并且 j-i < j-deque[last]
			// 所以，没有必要把 deque[last] 放入队列中
			last--
		}
		last++
		deque[last] = i
	}

	if res == initialValue {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
