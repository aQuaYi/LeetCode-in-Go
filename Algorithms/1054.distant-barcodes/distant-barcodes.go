package problem1054

func rearrangeBarcodes(A []int) []int {
	n := len(A)

	count := [10001]int{}
	max, maxA := 0, 0
	for _, a := range A {
		count[a]++
		if max < count[a] {
			max, maxA = count[a], a
		}
	}

	res := make([]int, n)
	for i := 0; i < max*2; i += 2 {
		res[i] = maxA
	}

	i := 1
	for _, a := range A {
		if a == maxA {
			continue
		}
		for res[i] != 0 {
			i++
		}
		res[i] = a
	}

	return res
}
