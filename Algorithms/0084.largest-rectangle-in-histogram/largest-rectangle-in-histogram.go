package Problem0084

func largestRectangleArea(heights []int) int {
	max := 0
	for i := 0; i < len(heights); i++ {
		lowest := heights[i]
		for j := i; j < len(heights); j++ {
			if lowest > heights[j] {
				lowest = heights[j]
			}

			temp := lowest * (j - i + 1)
			if max < temp {
				max = temp
			}
		}
	}

	return max
}
