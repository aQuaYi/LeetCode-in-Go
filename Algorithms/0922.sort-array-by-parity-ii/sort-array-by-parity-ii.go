package problem0922

func sortArrayByParityII(A []int) []int {
	size := len(A)
	res := make([]int, size)
	even, odd := 0, 1

	for _, a := range A {
		if a%2 == 0 {
			res[even] = a
			even += 2
		} else {
			res[odd] = a
			odd += 2
		}
	}

	return res
}
