package Problem0410

func splitArray(nums []int, m int) int {
	var Max, n, sum int
	for _, n = range nums {
		Max = max(Max, n)
		sum += n
	}

	if m == 1 {
		return sum
	}

	if Max > sum/m {
		return Max
	}

	valid := func(target, m int) bool {
		count, total := 1, 0
		for _, n = range nums {
			total += n
			if total > target {
				total = n
				count++
				if count > m {
					return false
				}
			}
		}
		return true
	}

	l, r := Max, sum
	var mid int
	for l <= r {
		mid = (l + r) >> 1
		if valid(mid, m) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
