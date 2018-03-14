package problem0209

func minSubArrayLen(s int, a []int) int {
	n := len(a)
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

	// res % (n+1) 是为了处理 sum(a) < s 的情况
	return res % (n + 1)
}
