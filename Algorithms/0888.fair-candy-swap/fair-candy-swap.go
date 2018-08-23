package problem0888

func fairCandySwap(A []int, B []int) []int {

	sumA := 0
	isExist := make(map[int]bool, len(A))
	for _, a := range A {
		sumA += a
		isExist[a] = true
	}

	sumB := 0
	for _, b := range B {
		sumB += b
	}

	halfDiff := (sumA - sumB) / 2

	a, b := 0, 0
	for _, b = range B {
		a = b + halfDiff
		if isExist[a] {
			break
		}
	}

	return []int{a, b}
}
