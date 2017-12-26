package Problem0713

func numSubarrayProductLessThanK(a []int, k int) int {
	res := 0

	n := len(a)
	p := 1
	i, j := 0, 0
	for i < n {
		for j < n && p*a[j] < k {
			p *= a[j]
			j++
		}

		if i == j {
			// 此时 a[i] > k
			// 需要跳过 a[i]
			// 由于 p 没有乘以 a[i]
			// 所以，p 也不需要除以 a[i]
			j++
		} else {
			res += j - i
			p /= a[i]
		}

		i++
	}

	return res
}
