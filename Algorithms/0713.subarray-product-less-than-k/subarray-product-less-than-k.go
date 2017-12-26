package Problem0713

func numSubarrayProductLessThanK(a []int, k int) int {
	res := 0

	n := len(a)
	p := 1
	i, j := 0, 0
	for i < n {
		for j < n && p < k {
			p *= a[j]
			j++
		}

		res += j - i - 1

		p /= a[i]
		i++
	}

	return res
}
