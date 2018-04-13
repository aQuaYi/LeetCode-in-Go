package problem0795

func numSubarrayBoundedMax(a []int, l int, r int) int {
	outOfRange := func(x int) bool {
		return x < l || r < x
	}
	res := 0
	lena := len(a)
	for i := 0; i < lena; i++ {
		if outOfRange(a[i]) {
			continue
		}
		left := i - 1
		for left >= 0 && a[left] < a[i] {
			left--
		}
		left = i - left - 1

		right := i
		for right < lena && a[right] <= a[i] {
			right++
		}
		right = right - i - 1

		res += 1 + left + right + left*right
	}

	return res
}
