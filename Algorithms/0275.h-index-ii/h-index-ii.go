package Problem0275

func hIndex(citations []int) int {
	size := len(citations)
	h := 0
	for h < size && citations[size-h-1] > h {
		h++
	}
	return h
}
