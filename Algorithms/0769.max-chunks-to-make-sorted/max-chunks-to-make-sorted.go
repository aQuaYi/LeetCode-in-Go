package Problem0769

func maxChunksToSorted(arr []int) int {
	lastIdx, res := 0, 0
	for i := 0; i < len(arr); i++ {
		if lastIdx < arr[i] {
			lastIdx = arr[i]
			continue
		}

		if i == lastIdx {
			res++
			lastIdx++
		}

	}

	return res
}
