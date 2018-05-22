package problem0031

func nextPermutation(a []int) {
	left := len(a) - 2
	for 0 <= left && a[left] >= a[left+1] {
		left--
	}

	// 此时 a[left+1:] 是一个 递减 数列

	reverse(a, left+1)

	if left == -1 {
		return
	}

	// 此时 a[left+1:] 是一个 递增 数列

	right := search(a, left+1, a[left])
	a[left], a[right] = a[right], a[left]
}

// 逆转 a[l:]
func reverse(a []int, l int) {
	r := len(a) - 1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}

// 返回 a[l:] 中 > target 的最小值的索引号
// a[l:] 是一个 递增 数列
func search(a []int, l, target int) int {
	r := len(a) - 1
	l--
	for l+1 < r {
		mid := (l + r) / 2
		if target < a[mid] {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}
