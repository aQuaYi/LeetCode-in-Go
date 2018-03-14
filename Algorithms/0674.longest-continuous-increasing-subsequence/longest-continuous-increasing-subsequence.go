package problem0674

func findLengthOfLCIS(a []int) int {
	if len(a) == 0 {
		return 0
	}

	res := 1

	i, j := 0, 1
	for j < len(a) {
		for j < len(a) && a[j-1] < a[j] {
			j++
		}

		if res < j-i {
			res = j - i
		}

		i = j
		j++
	}

	return res
}
