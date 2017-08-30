package Problem0565

func arrayNesting(a []int) int {
	max := 0
	for i := 0; i < len(a); i++ {
		temp := 1
		for a[i] != i {
			a[i], a[a[i]] = a[a[i]], a[i]
			temp++
		}
		if max < temp {
			max = temp
		}
	}
	return max
}
