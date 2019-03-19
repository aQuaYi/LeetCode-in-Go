package problem0961

func repeatedNTimes(A []int) int {
	hasSeen := [10000]bool{}
	var res int
	for _, res = range A {
		if hasSeen[res] {
			break
		}
		hasSeen[res] = true
	}
	return res
}
