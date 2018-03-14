package problem0667

func constructArray(n int, k int) []int {
	res := make([]int, 0, n)

	i, j := 1, n

	for i <= j {
		if k%2 == 1 {
			res = append(res, i)
			i++
		} else {
			res = append(res, j)
			j--
		}
		if k > 1 {
			k--
		}
	}

	return res
}
