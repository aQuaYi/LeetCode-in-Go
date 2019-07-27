package problem1122

func relativeSortArray(A, B []int) []int {
	count := [1001]int{}
	for _, a := range A {
		count[a]++
	}

	res := make([]int, 0, len(A))
	for _, b := range B {
		for count[b] > 0 {
			res = append(res, b)
			count[b]--
		}
	}
	for i := 0; i < 1001; i++ {
		for count[i] > 0 {
			res = append(res, i)
			count[i]--
		}
	}

	return res
}
