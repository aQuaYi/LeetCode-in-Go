package problem0410

func splitArray(nums []int, m int) int {
	var max, n, sum int
	for _, n = range nums {
		sum += n
		if max < n {
			max = n
		}
	}

	if m == 1 {
		return sum
	}

	// guess 是我们猜想的答案
	// isBigger 用来判断是否 guess >= 实际结果
	isBigger := func(guess int) bool {
		// nums 会被切分成连续的 子数组
		// 切分标准是，这段连续子数组的和是 不超过 guess 的最大值
		count, subSum := 1, 0
		for _, n = range nums {
			subSum += n
			if subSum > guess {
				subSum = n
				// n 已经是第 count+1 个子数组的成员了
				// 所以，count++
				count++
				if count > m {
					// 导致 nums 被切分的段数超过了 m
					// 可知 guess < 实际结果
					return false
				}
			}
		}
		// guess >= 实际结果
		// 才能导致 nums 被切分的段数，不超过 m
		return true
	}

	// 二分查找
	// 可知， max<= res <= sum
	// 那就在 max 和 sum 之间，进行二分查找
	l, r := max, sum
	var mid int
	for l <= r {
		mid = l + (r-l)>>1
		if isBigger(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}
