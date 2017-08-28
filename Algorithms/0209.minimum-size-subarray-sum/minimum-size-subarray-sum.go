package Problem0209

func minSubArrayLen(s int, a []int) int {
	n := len(a)

	if n == 0 {
		return 0
	}

	res, i, j, sum := n+1, 0, 0, 0

	for j < n {
		sum += a[j]
		j++

		for sum >= s {
			sum -= a[i]
			i++
			if res > j-i+1 {
				res = j - i + 1
			}
		}
	}

	return res % (n + 1)
}
