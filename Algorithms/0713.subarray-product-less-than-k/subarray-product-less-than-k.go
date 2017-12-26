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

		res += j - i

		p /= a[i]
		i++

		// if i > j {
		// 	j = i
		// }

	}

	return res
}
