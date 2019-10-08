package problem1207

func uniqueOccurrences(A []int) bool {
	count := make(map[int]int, len(A))
	for _, a := range A {
		count[a]++
	}
	hasSeen := make(map[int]bool, len(count))
	for _, c := range count {
		if hasSeen[c] {
			return false
		}
		hasSeen[c] = true
	}
	return true
}
