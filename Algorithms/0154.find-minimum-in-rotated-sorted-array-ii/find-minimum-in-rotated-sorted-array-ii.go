package problem0154

func findMin(a []int) int {
	L := len(a)

	if a[0] < a[L-1] {
		return a[0]
	}

	i, j := 0, L-1
	for i < j {
		if a[i] > a[i+1] {
			return a[i+1]
		}

		if a[j-1] > a[j] {
			return a[j]
		}

		i++
		j--
	}

	return a[i]
}
