package Problem0756

func pourWater(heights []int, V int, K int) []int {
	if V == 0 {
		return heights
	}

	leftBeg, leftEnd := K, K
	for 0 <= leftBeg-1 && heights[leftBeg-1] <= heights[leftBeg] {
		if heights[leftBeg-1] < heights[leftBeg] {
			leftEnd = leftBeg
		}
		leftBeg--
	}
	if heights[leftBeg] < heights[leftEnd] {
		for i := leftEnd - 1; leftBeg <= i; i-- {
			heights[i]++
			V--
			if V == 0 {
				return heights
			}
		}
		return pourWater(heights, V, K)
	}

	size := len(heights)
	rightBeg, rightEnd := K, K
	for rightBeg+1 < size && heights[rightBeg] >= heights[rightBeg+1] {
		if heights[rightBeg] > heights[rightBeg+1] {
			rightEnd = rightBeg
		}
		rightBeg++
	}
	if heights[rightBeg] < heights[rightEnd] {
		for i := rightEnd + 1; i <= rightBeg; i++ {
			heights[i]++
			V--
			if V == 0 {
				return heights
			}
		}
		return pourWater(heights, V, K)
	}

	for i := K; leftBeg <= i; i-- {
		heights[i]++
		V--
		if V == 0 {
			return heights
		}
	}

	for j := K + 1; j <= rightBeg; j++ {
		heights[j]++
		V--
		if V == 0 {
			return heights
		}
	}

	return pourWater(heights, V, K)
}
