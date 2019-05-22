package problem1013

func canThreePartsEqualSum(A []int) bool {
	s := sum(A)
	if s%3 != 0 {
		return false
	}

	return false
}

func sum(A []int) int {
	sum := 0
	for _, a := range A {
		sum += a
	}
	return sum
}
