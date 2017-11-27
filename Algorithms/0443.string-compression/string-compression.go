package Problem0443

func compress(chars []byte) int {
	r := [128]int{}
	for _, c := range chars {
		r[c]++
	}

	res := 0

	for i := 0; i < 128; i++ {
		if r[i] <= 1 {
			res += r[i]
			continue
		}

		for r[i] > 0 {
			res++
			r[i] /= 10
		}
		res++
	}

	return res
}
