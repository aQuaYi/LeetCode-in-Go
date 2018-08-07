package problem0862

func shortestSubarray(a []int, K int) int {
	size := len(a)

	sums := make([]int, size+1)
	for i, n := range a {
		if n == K {
			return 1
		}
		sums[i+1] = sums[i] + n
	}

	res := size + 1

	deque := make([]int, size+2)
	l, r := 0, -1

	for i := 0; i < size+1; i++ {
		for r-l >= 0 && sums[i]-sums[deque[l]] >= K {
			// 由于 i 递增
			// 也许以后会有 j>i 使得 sums[j] - sums[deque[l]] >= K
			// 但是由于 j>i, j-deque[l] > i-deque[l] 不会是更短的答案。
			// 所以，可以把 deque[l] 删除
			res = min(res, i-deque[l])
			l++
		}
		for r-l >= 0 && sums[i] <= sums[deque[r]] {
			// 如果存在 j>i>deque[r] 使得 sums[j] - sums[deque[r]] >= K
			// 由于 sums[deque[r]] >= sums[i] 则
			// sums[j] - sums[i] >= K 一定成立，并且 j-i < j-deque[r]
			// 所以，没有必要把 deque[r] 放入队列中
			r--
		}
		r++
		deque[r] = i
	}

	if res == size+1 {
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
