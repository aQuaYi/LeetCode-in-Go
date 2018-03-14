package problem0713

func numSubarrayProductLessThanK(a []int, k int) int {
	res := 0

	n := len(a)
	p := 1

	j := 0
	for i := 0; i < n; i++ {
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
	}

	return res
}
