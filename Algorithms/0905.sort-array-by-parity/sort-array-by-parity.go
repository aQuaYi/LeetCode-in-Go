package problem0905

func sortArrayByParity(a []int) []int {
	i, j, size := 0, len(a)-1, len(a)
	for i < j {
		for i < size && a[i]%2 == 0 {
			i++
		}
		for i < j && a[j]%2 == 1 {
			j--
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	return a
}
