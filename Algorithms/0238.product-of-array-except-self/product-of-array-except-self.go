package Problem0238

func productExceptSelf(a []int) []int {
	l := len(a)
	res := make([]int, l)

	p := 1
	zero := 0
	for i := 0; i < l; i++ {
		if a[i] == 0 {
			zero++
			continue
		}
		p *= a[i]
	}

	if zero > 1 {
		return res
	}

	for i := 0; i < l; i++ {
		if a[i] != 0 {
			res[i] = p / a[i]
		}
	}

	return res
}
