package Problem0443

func compress(chars []byte) int {
	res := 0
	i, j := 0, 0

	for i < len(chars) {
		res++
		for j < len(chars) && chars[i] == chars[j] {
			j++
		}

		t := j - i
		i = j

		if t == 1 {
			continue
		}

		for t > 0 {
			res++
			t /= 10
		}
	}

	return res
}
