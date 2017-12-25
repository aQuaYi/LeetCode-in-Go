package Problem0674

func findLengthOfLCIS(a []int) int {

	res := 1

	var i, j int
	i = 0
	j = 1

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
