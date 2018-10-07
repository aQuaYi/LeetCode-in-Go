package problem0905

func sortArrayByParity(a []int) []int {
	i, j := 0, len(a)-1
	for {
		for i < j && a[i]%2 == 0 {
			i++
		}
		for i < j && a[j]%2 == 1 {
			j--
		}

		if i < j {
			a[i], a[j] = a[j], a[i]
		} else {
			break
		}

	}
	return a
}
