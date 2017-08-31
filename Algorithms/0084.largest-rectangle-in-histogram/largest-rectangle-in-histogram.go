package Problem0084

func largestRectangleArea(heights []int) int {
	max := 0
	Len := len(heights)

	if Len == 0 {
		return 0
	}

	if Len == 1 {
		return heights[0]
	}

	min := func(first, last int) int {
		res := heights[first]
		for i := first + 1; i <= last; i++ {
			if res > heights[i] {
				res = heights[i]
			}
		}
		return res
	}

	i, j := 0, Len-1
	for i <= j {
		temp := (j - i + 1) * min(i, j)
		if max < temp {
			max = temp
		}
		if heights[i] < heights[j] ||
			(i+1 < Len && 0 <= j-1 && i+1 < j-1 && heights[i+1] < heights[j-1]) {
			i++
		} else {
			j--
		}
	}

	return max
}
