package problem0440

func findKthNumber(n int, k int) int {
	/*
		[1,n] 经过 lexicographical 排序后，其顺序集合 s
		for res:=1; res<=9; res++ {
			s += [res, res+1) +
				 [res*10, (res+1)*10 ) +
				 [res*100, (res+1)*100 ) +
				 ...
		}
	*/
	res := 1
	// k 与 res 总是同时更新
	// 当 k == 0 时， res 就是所求的答案
	k--
	for k > 0 {
		// count 统计了在 [1, n] 内，以 res 开头的数的个数
		count := 0
		begin, end := res, res+1
		for begin <= n {
			count += min(n+1, end) - begin
			begin *= 10
			end *= 10
		}

		if count <= k {
			// 答案不在以 res 开头的数中
			// 排除以 res 开头的 count 个数
			k -= count
			// 去下一个开头中查找
			res++
		} else {
			// 答案在以 res 开头的数中
			// 但是不在 [res, res+1) 中，
			// 因为在的话，k 在上一个 for 循环结束时，就是 0 了
			// 排除了 [res, res+1) 中只有一个数字
			k--
			// 在 res*10 + [0, 9] 这些开头中，重新开始查找
			res *= 10
		}

		// 每个 for 循环都可以
		//   在 if   中排除一部分范围
		//   在 else 中获取 res 的前缀的更多细节
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
