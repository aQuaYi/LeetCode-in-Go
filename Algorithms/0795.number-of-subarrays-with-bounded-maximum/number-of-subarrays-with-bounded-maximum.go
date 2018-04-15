package problem0795

func numSubarrayBoundedMax(a []int, l int, r int) int {
	res, heads, tails := 0, 0, 0
	isHead := func(x int) bool {
		return l <= x && x <= r
	}
	isTail := func(x int) bool {
		return x < l
	}

	// 每一个合格的 subarray 都必须至少含有一个 head

	// 在按顺序迭代的过程中
	// res += 以 a[i] 为末尾元素的 subarray 的个数，如果 a[i] 是 head 或 tail 的话
	for i := 0; i < len(a); i++ {
		if isTail(a[i]) {
			tails++
			// 如果 a[i] 是 tail 的话
			// 此前所有包含了 head 的 subarray 都能够延续到 a[i] 成为新的 subarray
			// 所以 res += heads，要把这些 heads 统计进去
			res += heads
		} else if isHead(a[i]) {
			// 如果 a[i] 是 head
			// heads = heads + tails + 1
			//           ↑       ↑     ↑
			//     	     |       |    a[i:i+1] 单独一个元素组成的 subarray
			//           |       a[i-x:i+1], 1<=x<=tails 这 tails 个 subarray
			//           a[i-y:i+1], tails < y <= heads + tails 这 heads 个 subarray
			// 可以看到这些 subarray 的末尾都是 a[i]
			heads += tails + 1
			// tails 重置为 0 是因为，算上 a[i] 后
			// a[i-z:i+1] 的末尾，不存在 tail
			tails = 0
			res += heads
		} else {
			// 超过 head 的话，需要重新计算
			heads = 0
			tails = 0
		}
	}

	return res
}
