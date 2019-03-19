package problem0961

func repeatedNTimes(A []int) int {
	N := len(A) / 2
	hasSeen := make(map[int]bool, N+1)
	var v int
	for _, v = range A {
		if hasSeen[v] {
			break
		}
		hasSeen[v] = true
	}
	return v
}
