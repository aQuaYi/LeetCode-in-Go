package Problem0263

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	if num <= 6 {
		return true
	}

	a := []int{2, 3, 5}
	for _, n := range a {
		for num%n == 0 {
			num /= n
		}
	}

	return num == 1
}
